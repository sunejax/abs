package config

type Model struct {
	Server struct {
		HTTP struct {
			Port string
		}
	}
	DataStores struct {
		Store struct {
			Auth        bool
			UseTLS      bool
			Database    string
			Hosts       []string
			User        string
			Password    string
			AuthSource  string
			Collections struct {
				Product string
				User    string
			}
		}
		Cart struct {
			Database    string
			User        string
			Password    string
			Hosts       []string
			Port        int
			Collections struct {
				Cart struct {
					Collection   string
					ModuleName   string
					FunctionName string
				}
			}
		}
	}
	Product struct {
		Inventory int64
	}
}
