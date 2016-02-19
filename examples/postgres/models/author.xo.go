// Package models contains the types for schema 'public'.
package models

// GENERATED BY XO. DO NOT EDIT.

import "errors"

// Author represents a row from public.authors.
type Author struct {
	AuthorID int    // author_id
	Name     string // name

	// xo fields
	_exists, _deleted bool
}

// Exists determines if the Author exists in the database.
func (a *Author) Exists() bool {
	return a._exists
}

// Deleted provides information if the Author has been deleted from the database.
func (a *Author) Deleted() bool {
	return a._deleted
}

// Insert inserts the Author to the database.
func (a *Author) Insert(db XODB) error {
	var err error

	// if already exist, bail
	if a._exists {
		return errors.New("insert failed: already exists")
	}

	// sql query
	const sqlstr = `INSERT INTO public.authors (` +
		`name` +
		`) VALUES (` +
		`$1` +
		`) RETURNING author_id`

	// run query
	XOLog(sqlstr, a.Name)
	err = db.QueryRow(sqlstr, a.Name).Scan(&a.AuthorID)
	if err != nil {
		return err
	}

	// set existence
	a._exists = true

	return nil
}

// Update updates the Author in the database.
func (a *Author) Update(db XODB) error {
	var err error

	// if doesn't exist, bail
	if !a._exists {
		return errors.New("update failed: does not exist")
	}

	// if deleted, bail
	if a._deleted {
		return errors.New("update failed: marked for deletion")
	}

	// sql query
	const sqlstr = `UPDATE public.authors SET (` +
		`name` +
		`) = ( ` +
		`$1` +
		`) WHERE author_id = $2`

	// run query
	XOLog(sqlstr, a.Name, a.AuthorID)
	_, err = db.Exec(sqlstr, a.Name, a.AuthorID)
	return err
}

// Save saves the Author to the database.
func (a *Author) Save(db XODB) error {
	if a.Exists() {
		return a.Update(db)
	}

	return a.Insert(db)
}

// Upsert performs an upsert for Author.
//
// NOTE: PostgreSQL 9.5+ only
func (a *Author) Upsert(db XODB) error {
	var err error

	// if already exist, bail
	if a._exists {
		return errors.New("insert failed: already exists")
	}

	// sql query
	const sqlstr = `INSERT INTO public.authors (` +
		`author_id, name` +
		`) VALUES (` +
		`$1, $2` +
		`) ON CONFLICT (author_id) DO UPDATE SET (` +
		`author_id, name` +
		`) = (` +
		`EXCLUDED.author_id, EXCLUDED.name` +
		`)`

	// run query
	XOLog(sqlstr, a.AuthorID, a.Name)
	_, err = db.Exec(sqlstr, a.AuthorID, a.Name)
	if err != nil {
		return err
	}

	// set existence
	a._exists = true

	return nil
}

// Delete deletes the Author from the database.
func (a *Author) Delete(db XODB) error {
	var err error

	// if doesn't exist, bail
	if !a._exists {
		return nil
	}

	// if deleted, bail
	if a._deleted {
		return nil
	}

	// sql query
	const sqlstr = `DELETE FROM public.authors WHERE author_id = $1`

	// run query
	XOLog(sqlstr, a.AuthorID)
	_, err = db.Exec(sqlstr, a.AuthorID)
	if err != nil {
		return err
	}

	// set deleted
	a._deleted = true

	return nil
}

// AuthorByAuthorID retrieves a row from public.authors as a Author.
//
// Looks up using index authors_pkey.
func AuthorByAuthorID(db XODB, authorID int) (*Author, error) {
	var err error

	// sql query
	const sqlstr = `SELECT ` +
		`author_id, name ` +
		`FROM public.authors ` +
		`WHERE author_id = $1`

	a := Author{
		_exists: true,
	}

	// run query
	XOLog(sqlstr, authorID)
	err = db.QueryRow(sqlstr, authorID).Scan(&a.AuthorID, &a.Name)
	if err != nil {
		return nil, err
	}

	return &a, nil
}

// AuthorsByName retrieves rows from public.authors, each as a Author.
//
// Looks up using index authors_name_idx.
func AuthorsByName(db XODB, name string) ([]*Author, error) {
	var err error

	// sql query
	const sqlstr = `SELECT ` +
		`author_id, name ` +
		`FROM public.authors ` +
		`WHERE name = $1`

	// run query
	XOLog(sqlstr, name)
	q, err := db.Query(sqlstr, name)
	if err != nil {
		return nil, err
	}
	defer q.Close()

	// load results
	res := []*Author{}
	for q.Next() {
		a := Author{
			_exists: true,
		}

		// scan
		err = q.Scan(&a.AuthorID, &a.Name)
		if err != nil {
			return nil, err
		}

		res = append(res, &a)
	}

	return res, nil
}