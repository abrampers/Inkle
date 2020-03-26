package logging

import (
	"testing"
	"time"
)

func isEventsEqual(a, b []*EventLog) bool {
	lena, lenb := len(a), len(b)

	if lena != lenb {
		return false
	}

	for i := 0; i < lena; i++ {
		if !isEventEqual(a[i], b[i]) {
			return false
		}
	}
	return true
}

func TestIsEventsEqual(t *testing.T) {
	tests := []struct {
		a    []*EventLog
		b    []*EventLog
		want bool
	}{
		{
			a:    []*EventLog{},
			b:    []*EventLog{},
			want: true,
		},
		{
			a: []*EventLog{},
			b: []*EventLog{
				&EventLog{},
			},
			want: false,
		},
		{
			a: []*EventLog{
				&EventLog{},
			},
			b:    []*EventLog{},
			want: false,
		},
		{
			a: []*EventLog{
				&EventLog{},
			},
			b: []*EventLog{
				&EventLog{},
			},
			want: true,
		},
		{
			a: []*EventLog{
				&EventLog{},
			},
			b: []*EventLog{
				&EventLog{ipsource: "::1"},
			},
			want: false,
		},
	}

	for _, test := range tests {
		if ret := isEventsEqual(test.a, test.b); ret != test.want {
			t.Errorf("isEventsEqual(a, b) returns '%t' while it should be '%t'", ret, test.want)
		}
	}
}

func TestCreateEvent(t *testing.T) {
	currtime := time.Now()
	tests := []struct {
		timestamp     time.Time
		servicename   string
		methodname    string
		ipsource      string
		tcpsource     string
		ipdest        string
		tcpdest       string
		initialevents []*EventLog
		finalevents   []*EventLog
	}{
		{
			timestamp:     currtime,
			servicename:   "helloworld.Greeter",
			methodname:    "SayHello",
			ipsource:      "::1",
			tcpsource:     "58108",
			ipdest:        "::1",
			tcpdest:       "8000",
			initialevents: []*EventLog{},
			finalevents: []*EventLog{
				&EventLog{
					tstart:      currtime,
					servicename: "helloworld.Greeter",
					methodname:  "SayHello",
					ipsource:    "::1",
					tcpsource:   "58108",
					ipdest:      "::1",
					tcpdest:     "8000",
					info:        "Request",
				},
			},
		},
		{
			timestamp:   currtime,
			servicename: "helloworld.Greeter",
			methodname:  "SayHello",
			ipsource:    "::1",
			tcpsource:   "58108",
			ipdest:      "::1",
			tcpdest:     "8000",
			initialevents: []*EventLog{
				&EventLog{},
			},
			finalevents: []*EventLog{
				&EventLog{},
				&EventLog{
					tstart:      currtime,
					servicename: "helloworld.Greeter",
					methodname:  "SayHello",
					ipsource:    "::1",
					tcpsource:   "58108",
					ipdest:      "::1",
					tcpdest:     "8000",
					info:        "Request",
				},
			},
		},
	}

	for _, test := range tests {
		elm := &eventLogManager{events: test.initialevents}
		elm.CreateEvent(test.timestamp, test.servicename, test.methodname, test.ipsource, test.tcpsource, test.ipdest, test.tcpdest)
		if !isEventsEqual(elm.events, test.finalevents) {
			t.Error("CreateEvent not working as expected")
		}
	}
}

// func TestInsertResponse(t *testing.T) {
// 	tests := []struct {
// 		initialevents []EventLog
// 		finalevents   []EventLog
// 	}{}
//
// 	for _, test := range tests {
// 	}
// }

func TestGetEvent(t *testing.T) {
	currtime := time.Now()
	tests := []struct {
		ipsource, tcpsource, ipdest, tcpdest string
		events                               []*EventLog
		idx                                  int
	}{
		{
			ipsource:  "::1",
			tcpsource: "58108",
			ipdest:    "::1",
			tcpdest:   "8000",
			events: []*EventLog{
				&EventLog{
					tstart:      currtime,
					servicename: "helloworld.Greeter",
					methodname:  "SayHello",
					ipsource:    "::1",
					tcpsource:   "58108",
					ipdest:      "::1",
					tcpdest:     "8000",
					info:        "Request",
				},
			},
			idx: 0,
		},
		{
			ipsource:  "127.0.0.1",
			tcpsource: "58108",
			ipdest:    "::1",
			tcpdest:   "8000",
			events: []*EventLog{
				&EventLog{
					tstart:      currtime,
					servicename: "helloworld.Greeter",
					methodname:  "SayHello",
					ipsource:    "::1",
					tcpsource:   "58108",
					ipdest:      "::1",
					tcpdest:     "8000",
					info:        "Request",
				},
			},
			idx: -1,
		},
		{
			ipsource:  "::1",
			tcpsource: "58100",
			ipdest:    "::1",
			tcpdest:   "8000",
			events: []*EventLog{
				&EventLog{
					tstart:      currtime,
					servicename: "helloworld.Greeter",
					methodname:  "SayHello",
					ipsource:    "::1",
					tcpsource:   "58108",
					ipdest:      "::1",
					tcpdest:     "8000",
					info:        "Request",
				},
			},
			idx: -1,
		},
		{
			ipsource:  "::1",
			tcpsource: "58108",
			ipdest:    "127.0.0.1",
			tcpdest:   "8000",
			events: []*EventLog{
				&EventLog{
					tstart:      currtime,
					servicename: "helloworld.Greeter",
					methodname:  "SayHello",
					ipsource:    "::1",
					tcpsource:   "58108",
					ipdest:      "::1",
					tcpdest:     "8000",
					info:        "Request",
				},
			},
			idx: -1,
		},
		{
			ipsource:  "::1",
			tcpsource: "58108",
			ipdest:    "::1",
			tcpdest:   "9000",
			events: []*EventLog{
				&EventLog{
					tstart:      currtime,
					servicename: "helloworld.Greeter",
					methodname:  "SayHello",
					ipsource:    "::1",
					tcpsource:   "58108",
					ipdest:      "::1",
					tcpdest:     "8000",
					info:        "Request",
				},
			},
			idx: -1,
		},
		{
			ipsource:  "::1",
			tcpsource: "58108",
			ipdest:    "::1",
			tcpdest:   "8000",
			events: []*EventLog{
				&EventLog{
					tstart:      currtime,
					servicename: "helloworld.Greeter",
					methodname:  "SayHello",
					ipsource:    "::1",
					tcpsource:   "58108",
					ipdest:      "::1",
					tcpdest:     "8000",
					info:        "Request",
				},
			},
			idx: 1,
		},
		{
			ipsource:  "127.0.0.1",
			tcpsource: "58108",
			ipdest:    "::1",
			tcpdest:   "8000",
			events: []*EventLog{
				&EventLog{},
				&EventLog{
					tstart:      currtime,
					servicename: "helloworld.Greeter",
					methodname:  "SayHello",
					ipsource:    "::1",
					tcpsource:   "58108",
					ipdest:      "::1",
					tcpdest:     "8000",
					info:        "Request",
				},
			},
			idx: -1,
		},
		{
			ipsource:  "::1",
			tcpsource: "58100",
			ipdest:    "::1",
			tcpdest:   "8000",
			events: []*EventLog{
				&EventLog{},
				&EventLog{
					tstart:      currtime,
					servicename: "helloworld.Greeter",
					methodname:  "SayHello",
					ipsource:    "::1",
					tcpsource:   "58108",
					ipdest:      "::1",
					tcpdest:     "8000",
					info:        "Request",
				},
			},
			idx: -1,
		},
		{
			ipsource:  "::1",
			tcpsource: "58108",
			ipdest:    "127.0.0.1",
			tcpdest:   "8000",
			events: []*EventLog{
				&EventLog{},
				&EventLog{
					tstart:      currtime,
					servicename: "helloworld.Greeter",
					methodname:  "SayHello",
					ipsource:    "::1",
					tcpsource:   "58108",
					ipdest:      "::1",
					tcpdest:     "8000",
					info:        "Request",
				},
			},
			idx: -1,
		},
		{
			ipsource:  "::1",
			tcpsource: "58108",
			ipdest:    "::1",
			tcpdest:   "9000",
			events: []*EventLog{
				&EventLog{},
				&EventLog{
					tstart:      currtime,
					servicename: "helloworld.Greeter",
					methodname:  "SayHello",
					ipsource:    "::1",
					tcpsource:   "58108",
					ipdest:      "::1",
					tcpdest:     "8000",
					info:        "Request",
				},
			},
			idx: -1,
		},
	}

	for _, test := range tests {
		elm := &eventLogManager{events: test.events}
		event, idx := elm.getEvent(test.ipsource, test.tcpsource, test.ipdest, test.tcpdest)
		if idx != test.idx {
			t.Errorf("getEvent returns wrong index. Expected '%d' got '%d'.", test.idx, idx)
		} else if idx != -1 && event != elm.events[idx] {
			t.Error("getEvent returns incorrect pointer.")
		}
	}
}

// func TestCleanupExpiredRequests(t *testing.T) {
// 	tests := []struct {
// 		initialevents []EventLog
// 		finalevents   []EventLog
// 	}{}
//
// 	for _, test := range tests {
// 	}
// }
