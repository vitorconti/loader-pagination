package seed

import (
	"bufio"
	"database/sql"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
	"sync"
)

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

func Seed(db *sql.DB) {
	source, err := os.Open("./music.csv")
	var wg sync.WaitGroup
	if err != nil {
		panic("error on opening file")
	}
	reader := bufio.NewReader(source)

	for {
		line, err := reader.ReadString('\n')
		if err != nil {
			break
		}
		line = strings.TrimSpace(line)
		var fields []interface{}
		for _, v := range strings.Split(line, ",") {
			fields = append(fields, strings.ReplaceAll(v, `"`, " "))
		}

		music := Music{
			ArtistFamiliarity:           parseStringToFloat64(fields[0].(string)),
			ArtistHotttnesss:            parseStringToFloat64(fields[1].(string)),
			ArtistID:                    fields[2].(string),
			ArtistLatitude:              parseStringToFloat64(fields[3].(string)),
			ArtistLocation:              parseStringToInt(fields[4].(string)),
			ArtistLongitude:             parseStringToFloat64(fields[5].(string)),
			ArtistName:                  fields[6].(string),
			ArtistSimilar:               parseStringToFloat64(fields[7].(string)),
			ArtistTerms:                 fields[8].(string),
			ArtistTermsFreq:             parseStringToFloat64(fields[9].(string)),
			ReleaseID:                   parseStringToInt(fields[10].(string)),
			ReleaseName:                 parseStringToInt(fields[11].(string)),
			SongArtistMbtags:            parseStringToFloat64(fields[12].(string)),
			SongArtistMbtagsCount:       parseStringToFloat64(fields[13].(string)),
			SongBarsConfidence:          parseStringToFloat64(fields[14].(string)),
			SongBarsStart:               parseStringToFloat64(fields[15].(string)),
			SongBeatsConfidence:         parseStringToFloat64(fields[16].(string)),
			SongBeatsStart:              parseStringToFloat64(fields[17].(string)),
			SongDuration:                parseStringToFloat64(fields[18].(string)),
			SongEndOfFadeIn:             parseStringToFloat64(fields[19].(string)),
			SongHotttnesss:              parseStringToFloat64(fields[20].(string)),
			SongID:                      fields[21].(string),
			SongKey:                     parseStringToInt(fields[22].(string)),
			SongKeyConfidence:           parseStringToFloat64(fields[23].(string)),
			SongLoudness:                parseStringToFloat64(fields[24].(string)),
			SongMode:                    parseStringToInt(fields[25].(string)),
			SongModeConfidence:          parseStringToFloat64(fields[26].(string)),
			SongStartOfFadeOut:          parseStringToFloat64(fields[27].(string)),
			SongTatumsConfidence:        parseStringToFloat64(fields[28].(string)),
			SongTatumsStart:             parseStringToFloat64(fields[29].(string)),
			SongTempo:                   parseStringToFloat64(fields[30].(string)),
			SongTimeSignature:           parseStringToInt(fields[31].(string)),
			SongTimeSignatureConfidence: parseStringToFloat64(fields[32].(string)),
			SongTitle:                   parseStringToInt(fields[33].(string)),
			SongYear:                    parseStringToInt(fields[34].(string)),
		}
		wg.Add(1)
		go populate(db, &music, &wg)
		wg.Wait()
	}

}
func parseStringToFloat64(stringToNumber string) float64 {
	conversion, err := strconv.ParseFloat(stringToNumber, 64)
	if err != nil {
		return conversion
	}
	fmt.Println(err)
	panic(err)
}
func parseStringToInt(stringToNumber string) int64 {
	conversion, err := strconv.ParseInt(stringToNumber, 16, 16)
	if err != nil {
		return conversion
	}
	fmt.Println(err)
	panic(err)
}
func populate(db *sql.DB, music *Music, wg *sync.WaitGroup) {
	defer wg.Done()
	query := `
        INSERT INTO music (
            artist_familiarity,
			artist_hotttnesss,
			artist_id,
			artist_latitude,
			artist_location,
			artist_longitude,
            artist_name,
			artist_similar,
			artist_terms,
			artist_terms_freq,
			release_id,
			release_name,
			song_artist_mbtags,
            song_artist_mbtags_count,
			song_bars_confidence,
			song_bars_start,
			song_beats_confidence,
			song_beats_start,
            song_duration,
			song_end_of_fade_in,
			song_hotttnesss,
			song_id,
			song_key,
			song_key_confidence,
			song_loudness,
            song_mode,
			song_mode_confidence,
			song_start_of_fade_out,
			song_tatums_confidence,
			song_tatums_start,
			song_tempo,
            song_time_signature,
			song_time_signature_confidence,
			song_title,
			song_year
        ) VALUES (?, ?, ?, ?, ?,?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)
    `

	_, err := db.Exec(
		query,
		music.ArtistFamiliarity,
		music.ArtistHotttnesss,
		music.ArtistID,
		music.ArtistLatitude,
		music.ArtistLocation,
		music.ArtistLongitude,
		music.ArtistName,
		music.ArtistSimilar,
		music.ArtistTerms,
		music.ArtistTermsFreq,
		music.ReleaseID,
		music.ReleaseName,
		music.SongArtistMbtags,
		music.SongArtistMbtagsCount,
		music.SongBarsConfidence,
		music.SongBarsStart,
		music.SongBeatsConfidence,
		music.SongBeatsStart,
		music.SongDuration,
		music.SongEndOfFadeIn,
		music.SongHotttnesss,
		music.SongID,
		music.SongKey,
		music.SongKeyConfidence,
		music.SongLoudness,
		music.SongMode,
		music.SongModeConfidence,
		music.SongStartOfFadeOut,
		music.SongTatumsConfidence,
		music.SongTatumsStart,
		music.SongTempo,
		music.SongTimeSignature,
		music.SongTimeSignatureConfidence,
		music.SongTitle,
		music.SongYear,
	)

	if err != nil {
		log.Fatal("Error executing insert query:", err)
	}

}
