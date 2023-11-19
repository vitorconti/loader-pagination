package entity

type Music struct {
	ArtistFamiliarity           float64 `json:"artistFamiliarity"`
	ArtistHotttnesss            float64 `json:"artistHotttnesss"`
	ArtistID                    string  `json:"artistID"`
	ArtistLatitude              float64 `json:"artistLatitude"`
	ArtistLocation              int64   `json:"artistLocation"`
	ArtistLongitude             float64 `json:"artistLongitude"`
	ArtistName                  string  `json:"artistName"`
	ArtistSimilar               float64 `json:"artistSimilar"`
	ArtistTerms                 string  `json:"artistTerms"`
	ArtistTermsFreq             float64 `json:"artistTermsFreq"`
	ReleaseID                   int64   `json:"releaseID"`
	ReleaseName                 int64   `json:"releaseName"`
	SongArtistMbtags            float64 `json:"songArtistMbtags"`
	SongArtistMbtagsCount       float64 `json:"songArtistMbtagsCount"`
	SongBarsConfidence          float64 `json:"songBarsConfidence"`
	SongBarsStart               float64 `json:"songBarsStart"`
	SongBeatsConfidence         float64 `json:"songBeatsConfidence"`
	SongBeatsStart              float64 `json:"songBeatsStart"`
	SongDuration                float64 `json:"songDuration"`
	SongEndOfFadeIn             float64 `json:"songEndOfFadeIn"`
	SongHotttnesss              float64 `json:"songHotttnesss"`
	SongID                      string  `json:"songID"`
	SongKey                     int64   `json:"songKey"`
	SongKeyConfidence           float64 `json:"songKeyConfidence"`
	SongLoudness                float64 `json:"songLoudness"`
	SongMode                    int64   `json:"songMode"`
	SongModeConfidence          float64 `json:"songModeConfidence"`
	SongStartOfFadeOut          float64 `json:"songStartOfFadeOut"`
	SongTatumsConfidence        float64 `json:"songTatumsConfidence"`
	SongTatumsStart             float64 `json:"songTatumsStart"`
	SongTempo                   float64 `json:"songTempo"`
	SongTimeSignature           int64   `json:"songTimeSignature"`
	SongTimeSignatureConfidence float64 `json:"songTimeSignatureConfidence"`
	SongTitle                   int64   `json:"songTitle"`
	SongYear                    int64   `json:"songYear"`
}
