package repositories

import (
	"database/sql"

	"github.com/flickr/entities"
)

type FlickrRepository struct {
	db *sql.DB
}

func NewFlickrRepo(db *sql.DB) *FlickrRepository {
	return &FlickrRepository{
		db: db,
	}
}

func (fr *FlickrRepository) OperateImage(image *entities.FlickrImage) error {
	stmt := `INSERT INTO images(title, link, author_id, vote) VALUES (?, ?,  ?,?)`
	statement, err := fr.db.Prepare(stmt)
	if err != nil {
		return err
	}
	_, err = statement.Exec(image.Title, image.Link, image.AuthorID, image.Vote)
	if err != nil {
		return err
	}
	return nil
}

func (fr *FlickrRepository) GetImage(image *entities.FlickrImage) (*entities.FlickrImage, error) {
	var img entities.FlickrImage
	row := fr.db.QueryRow("select * from images where title = $1 and link = $2 and author_id = $3 ",
		image.Title, image.Link, image.AuthorID)
	err := row.Scan(&img.ID, &img.Title, &img.Link, &img.AuthorID, &img.Vote)
	if err != nil && err != sql.ErrNoRows {
		return nil, err
	}
	return &img, nil
}

func (fr *FlickrRepository) UpdateVote(image *entities.FlickrImage) error {
	_, err := fr.db.Exec("update images set vote = $1 where title = $2 and link = $3 and author_id = $4", image.Vote, image.Title, image.Link, image.AuthorID)
	if err != nil {
		return err
	}
	return nil
}

func (fr *FlickrRepository) GetAllImages() ([]entities.FlickrImage, error) {
	images := []entities.FlickrImage{}
	rows, err := fr.db.Query("select * from images")
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		img := entities.FlickrImage{}
		if err := rows.Scan(
			&img.ID,
			&img.Title,
			&img.Link,
			&img.AuthorID,
			&img.Vote,
		); err != nil {
			return nil, err
		}
		images = append(images, img)
	}

	return images, nil
}
