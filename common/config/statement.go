package config

type (
	ConfigStatement struct {
		Database Database `json:"database"`
	}

	Database struct {
		Default DBItem `json:"default"`
	}

	DBItem struct {
		Db          string     `json:"db"`
		Password    string     `json:"password"`
		User        string     `json:"user"`
		Connections Connection `json:"connections"`
	}

	Connection struct {
		Master string   `json:"master"`
		Slaves []string `json:"slaves"`
	}
)
