package dto

type IndexSettings struct {
	Settings struct {
		Analysis struct {
			Analyzer struct {
				RussianAnalyzer struct {
					Type      string   `json:"type"`
					Tokenizer string   `json:"tokenizer"`
					Filter    []string `json:"filter"`
				} `json:"russian_analyzer"`
				SpellAnalyzer struct {
					Type      string   `json:"type"`
					Tokenizer string   `json:"tokenizer"`
					Filter    []string `json:"filter"`
				} `json:"spell_analyzer"`
			} `json:"analyzer"`
			Filter struct {
				SnowFilter struct {
					Type     string `json:"type"`
					Language string `json:"language"`
				} `json:"snow_filter"`
				NgramFilter struct {
					Type    string `json:"type"`
					MinGram int    `json:"min_gram"`
					MaxGram int    `json:"max_gram"`
				} `json:"ngram_filter"`
				SpellFilter struct {
					Type   string `json:"type"`
					Locale string `json:"locale"`
				} `json:"spell_filter"`
			} `json:"filter"`
		} `json:"analysis"`
	} `json:"settings"`
	Mappings struct {
		Properties struct {
			Title struct {
				Type     string `json:"type"`
				Analyzer string `json:"analyzer"`
				Fields   struct {
					Russian struct {
						Type           string `json:"type"`
						Analyzer       string `json:"analyzer"`
						SearchAnalyzer string `json:"search_analyzer"`
					} `json:"russian"`
					Spell struct {
						Type           string `json:"type"`
						Analyzer       string `json:"analyzer"`
						SearchAnalyzer string `json:"search_analyzer"`
					} `json:"spell"`
				} `json:"fields"`
			} `json:"title"`
			Description struct {
				Type     string `json:"type"`
				Analyzer string `json:"analyzer"`
				Fields   struct {
					Russian struct {
						Type           string `json:"type"`
						Analyzer       string `json:"analyzer"`
						SearchAnalyzer string `json:"search_analyzer"`
					} `json:"russian"`
					Spell struct {
						Type           string `json:"type"`
						Analyzer       string `json:"analyzer"`
						SearchAnalyzer string `json:"search_analyzer"`
					} `json:"spell"`
				} `json:"fields"`
			} `json:"description"`
		} `json:"properties"`
	} `json:"mappings"`
}

func InitSettings() IndexSettings {
	var s IndexSettings

	s.Settings.Analysis.Analyzer.RussianAnalyzer.Type = "custom"
	s.Settings.Analysis.Analyzer.RussianAnalyzer.Tokenizer = "standard"
	s.Settings.Analysis.Analyzer.RussianAnalyzer.Filter = []string{"lowercase", "snow_filter", "ngram_filter"}

	s.Settings.Analysis.Analyzer.SpellAnalyzer.Type = "custom"
	s.Settings.Analysis.Analyzer.SpellAnalyzer.Tokenizer = "standard"
	s.Settings.Analysis.Analyzer.SpellAnalyzer.Filter = []string{"lowercase", "spell_filter"}

	s.Settings.Analysis.Filter.SnowFilter.Type = "snowball"
	s.Settings.Analysis.Filter.SnowFilter.Language = "Russian"

	s.Settings.Analysis.Filter.NgramFilter.Type = "ngram"
	s.Settings.Analysis.Filter.NgramFilter.MinGram = 5
	s.Settings.Analysis.Filter.NgramFilter.MaxGram = 6

	s.Settings.Analysis.Filter.SpellFilter.Type = "hunspell"
	s.Settings.Analysis.Filter.SpellFilter.Locale = "ru_RU"

	s.Mappings.Properties.Title.Type = "text"
	s.Mappings.Properties.Title.Analyzer = "standard"
	s.Mappings.Properties.Title.Fields.Russian.Type = "text"
	s.Mappings.Properties.Title.Fields.Russian.Analyzer = "russian_analyzer"
	s.Mappings.Properties.Title.Fields.Russian.SearchAnalyzer = "russian_analyzer"
	s.Mappings.Properties.Title.Fields.Spell.Type = "text"
	s.Mappings.Properties.Title.Fields.Spell.Analyzer = "spell_analyzer"
	s.Mappings.Properties.Title.Fields.Spell.SearchAnalyzer = "spell_analyzer"

	s.Mappings.Properties.Description.Type = "text"
	s.Mappings.Properties.Description.Analyzer = "standard"
	s.Mappings.Properties.Description.Fields.Russian.Type = "text"
	s.Mappings.Properties.Description.Fields.Russian.Analyzer = "russian_analyzer"
	s.Mappings.Properties.Description.Fields.Russian.SearchAnalyzer = "russian_analyzer"
	s.Mappings.Properties.Description.Fields.Spell.Type = "text"
	s.Mappings.Properties.Description.Fields.Spell.Analyzer = "spell_analyzer"
	s.Mappings.Properties.Description.Fields.Spell.SearchAnalyzer = "spell_analyzer"

	return s
}
