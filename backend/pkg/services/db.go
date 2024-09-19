package services

import (
	"database/sql"
	"log"
	"os"
	"time"

	_ "github.com/mattn/go-sqlite3"
)

var db *sql.DB

// Duration a file exists for before it's automatically deleted
var expiryDuration time.Duration = time.Hour * 24

func InitDB() {
	var err error
	cwd, _ := os.Getwd()

	db, err = sql.Open("sqlite3", cwd+"/fileinfo.db")
	if err != nil {
		log.Fatal(err)
	}

	_, err = db.Exec(`
		create table if not exists video_downloads (
			id integer primary key autoincrement,
			video_id string,
			video_format_id string,
			audio_format_id string,
			unique_filename string,
			created_at datetime default current_timestamp
		)
	`)

	if err != nil {
		log.Fatal(err)
	}
}

func InsertDownload(videoID string, uniqueFilename string, videoFormatId string, audioFormatId string) (string, error) {
	_, err := db.Exec(`
		insert into video_downloads (video_id, unique_filename, video_format_id, audio_format_id)
		values (?, ?, ?, ?)`,
		videoID,
		uniqueFilename,
		videoFormatId,
		audioFormatId,
	)

	if err != nil {
		return "", err
	}
	return "Download inserted", err
}

func DeleteAllExpiredFiles() {
	expiryDateTime := time.Now().Add(expiryDuration)

	// TODO: we need to add the logic for deleting files

	_, err := db.Exec(`
		delete from video_downloads where created_at < ?`,
		expiryDateTime,
	)

	if err != nil {
		log.Fatal(err)
	}
}

