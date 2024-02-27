package tests

import (
	"database/sql"
	"database/sql/driver"
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/magiconair/properties/assert"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"movie-api/internal/domain"
	"movie-api/internal/model"
	"movie-api/internal/repository"
	"regexp"
	"testing"
	"time"
)

type Suite struct {
	suite.Suite
	DB   *gorm.DB
	mock sqlmock.Sqlmock

	repo  repository.MovieRepository
	movie *model.Movie
}

func (s *Suite) AfterTest(_, _ string) {
	require.NoError(s.T(), s.mock.ExpectationsWereMet())
}

func TestInit(t *testing.T) {
	suite.Run(t, new(Suite))
}

func (s *Suite) SetupSuite() {
	var (
		db  *sql.DB
		err error
	)

	db, s.mock, err = sqlmock.New()
	require.NoError(s.T(), err)

	s.DB, err = gorm.Open(postgres.New(postgres.Config{
		Conn: db,
	}), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})
	require.NoError(s.T(), err)

	s.repo = repository.NewMovieRepository(s.DB)
}

func (s *Suite) TestMovieRepo_1List() {
	timeNow := time.Now()
	values := [][]driver.Value{
		{1, "Agak Laen", "Film Agak Laen", 9.5, "agak_laen.jpg", false, timeNow, time.Time{}},
		{2, "Warkop DKI", "Warkop DKI", 9.5, "warkop.jpg", false, timeNow, time.Time{}}}

	expectedModel := &[]model.Movie{
		{
			ID:          1,
			Title:       "Agak Laen",
			Description: "Film Agak Laen",
			Rating:      9.5,
			Image:       "agak_laen.jpg",
			Deleted:     false,
			CreatedAt:   timeNow,
			UpdatedAt:   time.Time{},
		},
		{
			ID:          2,
			Title:       "Warkop DKI",
			Description: "Warkop DKI",
			Rating:      9.5,
			Image:       "warkop.jpg",
			Deleted:     false,
			CreatedAt:   timeNow,
			UpdatedAt:   time.Time{},
		},
	}

	s.mock.ExpectQuery(regexp.QuoteMeta(
		`SELECT * FROM "movies" WHERE deleted = $1`)).WillReturnRows(sqlmock.NewRows(
		[]string{"id", "title", "description", "rating", "image", "deleted", "created_at", "updated_at"}).AddRows(values...))

	results, err := s.repo.List()
	require.NoError(s.T(), err)
	assert.Equal(s.T(), results, expectedModel)
}

func (s *Suite) TestMovieRepo_2View() {
	timeNow := time.Now()

	expectedModel := &model.Movie{
		ID:          1,
		Title:       "Agak Laen",
		Description: "Film Agak Laen",
		Rating:      9.5,
		Image:       "agak_laen.jpg",
		Deleted:     false,
		CreatedAt:   timeNow,
		UpdatedAt:   time.Time{},
	}

	s.mock.ExpectQuery(regexp.QuoteMeta(
		`SELECT * FROM "movies" WHERE deleted=false AND "movies"."id" = $1 ORDER BY "movies"."id" LIMIT $2`)).WillReturnRows(sqlmock.NewRows(
		[]string{"id", "title", "description", "rating", "image", "deleted", "created_at", "updated_at"}).
		AddRow(1, "Agak Laen", "Film Agak Laen", 9.5, "agak_laen.jpg", false, timeNow, time.Time{}))

	results, err := s.repo.View(expectedModel.ID)
	require.NoError(s.T(), err)
	assert.Equal(s.T(), results, expectedModel)
}

func (s *Suite) TestMovieRepo_3Insert() {
	timeNow := time.Now()
	id := 1

	inputRequest := &domain.MovieRequest{
		ID:          uint(id),
		Title:       "Warkop DKI",
		Description: "Film Warkop DKI",
		Rating:      9.5,
		Image:       "warkop_dki.jpg",
	}

	expectedModel := &model.Movie{
		ID:          uint(id),
		Title:       "Warkop DKI",
		Description: "Film Warkop DKI",
		Rating:      9.5,
		Image:       "warkop_dki.jpg",
		Deleted:     false,
		CreatedAt:   timeNow,
		UpdatedAt:   timeNow,
	}

	s.mock.ExpectBegin()
	s.mock.ExpectQuery(regexp.QuoteMeta(
		`INSERT INTO "movies" ("title","description","rating","image","created_at","updated_at") VALUES ($1,$2,$3,$4,$5,$6) RETURNING "id","deleted"`)).
		WithArgs("Warkop DKI", "Film Warkop DKI", 9.5, "warkop_dki.jpg", timeNow, timeNow).
		WillReturnRows(sqlmock.NewRows(
			[]string{"id", "title", "description", "rating", "image", "created_at", "updated_at"}).
			AddRow(id, "Warkop DKI", "Film Warkop DKI", 9.5, "warkop_dki.jpg", timeNow, timeNow))
	s.mock.ExpectCommit()

	result, err := s.repo.Insert(inputRequest, timeNow)
	require.NoError(s.T(), err)
	assert.Equal(s.T(), result, expectedModel)
}
