package pokeapi

type RespLocations struct {
	Count    int    `json:"count"`
	Next     string `json:"next"`
	Previous string `json:"previous"`
	Results  []struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	} `json:"results"`
}

// type LocationAreaDetails struct {
// 	EncounterMethodRates []struct {
// 		EncounterMethod struct {
// 			Name string `json:"name"`
// 			URL  string `json:"url"`
// 		} `json:"encounter_method"`
// 		VersionDetails []struct {
// 			Rate    int `json:"rate"`
// 			Version struct {
// 				Name string `json:"name"`
// 				URL  string `json:"url"`
// 			} `json:"version"`
// 		} `json:"version_details"`
// 	} `json:"encounter_method_rates"`
// 	GameIndex int `json:"game_index"`
// 	ID        int `json:"id"`
// 	Location  struct {
// 		Name string `json:"name"`
// 		URL  string `json:"url"`
// 	} `json:"location"`
// 	Name  string `json:"name"`
// 	Names []struct {
// 		Language struct {
// 			Name string `json:"name"`
// 			URL  string `json:"url"`
// 		} `json:"language"`
// 		Name string `json:"name"`
// 	} `json:"names"`
// 	PokemonEncounters []struct {
// 		Pokemon struct {
// 			Name string `json:"name"`
// 			URL  string `json:"url"`
// 		} `json:"pokemon"`
// 		VersionDetails []struct {
// 			EncounterDetails []struct {
// 				Chance          int   `json:"chance"`
// 				ConditionValues []any `json:"condition_values"`
// 				MaxLevel        int   `json:"max_level"`
// 				Method          struct {
// 					Name string `json:"name"`
// 					URL  string `json:"url"`
// 				} `json:"method"`
// 				MinLevel int `json:"min_level"`
// 			} `json:"encounter_details"`
// 			MaxChance int `json:"max_chance"`
// 			Version   struct {
// 				Name string `json:"name"`
// 				URL  string `json:"url"`
// 			} `json:"version"`
// 		} `json:"version_details"`
// 	} `json:"pokemon_encounters"`
// }

type LocationAreaDetails struct {
	EncounterMethodRates []EncounterMethodRates `json:"encounter_method_rates"`
	GameIndex            int                    `json:"game_index"`
	ID                   int                    `json:"id"`
	Location             Location               `json:"location"`
	Name                 string                 `json:"name"`
	Names                []Names                `json:"names"`
	PokemonEncounters    []PokemonEncounters    `json:"pokemon_encounters"`
}
type EncounterMethod struct {
	Name string `json:"name"`
	URL  string `json:"url"`
}
type Version struct {
	Name string `json:"name"`
	URL  string `json:"url"`
}
type VersionDetails struct {
	Rate    int     `json:"rate"`
	Version Version `json:"version"`
}
type EncounterMethodRates struct {
	EncounterMethod EncounterMethod  `json:"encounter_method"`
	VersionDetails  []VersionDetails `json:"version_details"`
}
type Location struct {
	Name string `json:"name"`
	URL  string `json:"url"`
}
type Language struct {
	Name string `json:"name"`
	URL  string `json:"url"`
}
type Names struct {
	Language Language `json:"language"`
	Name     string   `json:"name"`
}
type Pokemon struct {
	Name string `json:"name"`
	URL  string `json:"url"`
}
type Method struct {
	Name string `json:"name"`
	URL  string `json:"url"`
}
type EncounterDetails struct {
	Chance          int    `json:"chance"`
	ConditionValues []any  `json:"condition_values"`
	MaxLevel        int    `json:"max_level"`
	Method          Method `json:"method"`
	MinLevel        int    `json:"min_level"`
}
type EncounterVersionDetails struct {
	EncounterDetails []EncounterDetails `json:"encounter_details"`
	MaxChance        int                `json:"max_chance"`
	Version          Version            `json:"version"`
}
type PokemonEncounters struct {
	Pokemon        Pokemon                   `json:"pokemon"`
	VersionDetails []EncounterVersionDetails `json:"version_details"`
}
