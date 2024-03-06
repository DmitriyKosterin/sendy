package migrations

import (
	"github.com/go-pg/migrations/v8"
)

func init() {
	migrations.MustRegisterTx(func(db migrations.DB) error {
		_, err := db.Exec(`
            CREATE TABLE publishers (
                publisher_id SERIAL PRIMARY KEY,
                name VARCHAR(255) NOT NULL
            );

            CREATE TABLE authors (
                author_id SERIAL PRIMARY KEY,
                name VARCHAR(255) NOT NULL
            );

            CREATE TABLE books (
                book_id SERIAL PRIMARY KEY,
                title VARCHAR(255) NOT NULL,
                publisher_id INTEGER REFERENCES publishers(publisher_id),
                author_id INTEGER REFERENCES authors(author_id)
            );

            INSERT INTO publishers (name) VALUES ('Издательство 2'), ('Издательство 3');

            INSERT INTO authors (name) VALUES 
            ('Автор 1'), ('Автор 3'), ('Автор 4'), ('Автор 6'), ('Автор 8'), ('Автор 9');

            INSERT INTO books (title, publisher_id, author_id) VALUES 
            ('Книга 1-1', null, 1), ('Книга 1-2', null, 1),
            ('Книга 2-1', null, null), ('Книга 2-2', null, null),
            ('Книга 3-1', null, 2), ('Книга 3-2', null, 2),
            ('Книга 4-1', 1, 3), ('Книга 4-2', 1, 3),
            ('Книга 5-1', 1, null), ('Книга 5-2', 1, null),
            ('Книга 6-1', 1, 4), ('Книга 6-2', 1, 4),
            ('Книга 7-1', 2, null), ('Книга 7-2', 2, null),
            ('Книга 8-1', 2, 5), ('Книга 8-2', 2, 5),
            ('Книга 9-1', 2, 6), ('Книга 9-2', 2, 6);
        `)
		return err
	}, func(db migrations.DB) error {
		_, err := db.Exec(`
            DROP TABLE IF EXISTS books;
            DROP TABLE IF EXISTS authors;
            DROP TABLE IF EXISTS publishers;
        `)
		return err
	})
}
