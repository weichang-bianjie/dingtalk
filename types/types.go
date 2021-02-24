package types

import "time"

type (
	DingTalkReq struct {
		Msgtype string  `json:"msgtype"`
		Text    Content `json:"text"`
	}
	Content struct {
		Content string `json:"content"`
	}

	AlertsContent struct {
		Receiver          string      `json:"receiver,omitempty"`
		Status            string      `json:"status,omitempty"`
		Alerts            []Alert     `json:"alerts,omitempty"`
		GroupLabels       interface{} `json:"groupLabels,omitempty"`
		CommonLabels      interface{} `json:"commonLabels,omitempty"`
		CommonAnnotations Annotation  `json:"commonAnnotations,omitempty"`
		ExternalURL       string      `json:"externalURL,omitempty"`
		Version           string      `json:"version,omitempty"`
		GroupKey          string      `json:"groupKey,omitempty"`
		TruncatedAlerts   int         `json:"truncatedAlerts,omitempty"`
	}
	Alert struct {
		Status       string      `json:"status,omitempty"`
		Labels       interface{} `json:"labels,omitempty"`
		Annotations  Annotation  `json:"annotations,omitempty"`
		StartsAt     time.Time   `json:"startsAt,omitempty"`
		EndsAt       time.Time   `json:"endsAt,omitempty"`
		GeneratorURL string      `json:"generatorURL,omitempty"`
		Fingerprint  string      `json:"fingerprint,omitempty"`
	}
	//Label struct {
	//	Alertname   string `json:"alertname,omitempty"`
	//	ExportedJob string `json:"exported_job,omitempty"`
	//	Instance    string `json:"instance,omitempty"`
	//	Job         string `json:"job,omitempty"`
	//	Name        string `json:"name,omitempty"`
	//	Sex         string `json:"sex,omitempty"`
	//}
	Annotation struct {
		Description string `json:"description,omitempty"`
		Summary     string `json:"summary,omitempty"`
	}
)
