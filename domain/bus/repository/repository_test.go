package repository

import (
	"testing"

	"errors"

	"regexp"

	"github.com/bouk/monkey"
	"github.com/rodrigo-brito/bus-api-go/domain/schedule/model"
	"github.com/rodrigo-brito/bus-api-go/domain/schedule/repository"
	"github.com/rodrigo-brito/bus-api-go/test/mysql"
	"github.com/smartystreets/assertions/should"
	. "github.com/smartystreets/goconvey/convey"
	"gopkg.in/DATA-DOG/go-sqlmock.v1"
)

func TestGetAll(t *testing.T) {
	Convey("Given a working database connection", t, func() {
		mock, err := mysql.GetDatabaseMock()
		defer mysql.UnmockDatabase()
		So(err, ShouldBeNil)
		Convey("When everything is OK", func() {
			Convey("It should execute the correct query and return a valid result", func() {
				query := queries["all"]
				So(query, should.NotBeEmpty)
				mock.ExpectQuery(query).WillReturnRows(
					sqlmock.NewRows([]string{"id", "number", "name", "fare"}).
						AddRow(1, 4988, "Bus One", 3.1))
				result, err := GetAll()
				So(err, ShouldBeNil)
				So(result, ShouldHaveLength, 1)
				So(result[0].ID, ShouldEqual, 1)
				So(*result[0].Number, ShouldEqual, "4988")
				So(result[0].Name, ShouldEqual, "Bus One")
				So(result[0].Fare, ShouldEqual, 3.1)
				So(mock.ExpectationsWereMet(), ShouldBeNil)
			})
		})
		Convey("When database fail", func() {
			Convey("It should execute the correct query and return error", func() {
				query := queries["all"]
				So(query, should.NotBeEmpty)

				expected := errors.New("fail")
				mock.ExpectQuery(query).WillReturnError(expected)
				result, err := GetAll()

				So(err, ShouldEqual, expected)
				So(result, ShouldBeNil)
				So(mock.ExpectationsWereMet(), ShouldBeNil)
			})
		})
	})
}

func TestGet(t *testing.T) {
	Convey("Given a working database connection", t, func() {
		mock, err := mysql.GetDatabaseMock()
		defer mysql.UnmockDatabase()
		So(err, ShouldBeNil)
		Convey("When everything is OK", func() {
			Convey("When the schedules is not required", func() {
				query := queries["by-id"]
				So(query, ShouldNotBeEmpty)

				mock.ExpectQuery(regexp.QuoteMeta(query)).WithArgs(int64(1)).WillReturnRows(
					sqlmock.NewRows([]string{"id", "number", "name", "fare"}).
						AddRow(1, "4988", "Bus One", 3.1))
				result, err := Get(1, false)

				So(err, ShouldBeNil)
				So(result.ID, ShouldEqual, 1)
				So(*result.Number, ShouldEqual, "4988")
				So(result.Name, ShouldEqual, "Bus One")
				So(result.Fare, ShouldEqual, 3.1)
				So(mock.ExpectationsWereMet(), ShouldBeNil)
			})
			Convey("When the schedules is required", func() {
				expectedSchedules := []*model.Schedule{
					new(model.Schedule),
					new(model.Schedule),
					new(model.Schedule),
				}
				monkey.Patch(repository.FetchManyByBus, func(busID int64) ([]*model.Schedule, error) {
					return expectedSchedules, nil
				})
				defer monkey.Unpatch(repository.FetchManyByBus)

				query := queries["by-id"]
				So(query, ShouldNotBeEmpty)

				mock.ExpectQuery(regexp.QuoteMeta(query)).WithArgs(int64(1)).WillReturnRows(
					sqlmock.NewRows([]string{"id", "number", "name", "fare"}).
						AddRow(1, "4988", "Bus One", 3.1))
				result, err := Get(1, true)

				So(err, ShouldBeNil)
				So(mock.ExpectationsWereMet(), ShouldBeNil)
				Convey("Schedules must be injected", func() {
					So(result.Schedules, ShouldHaveLength, len(expectedSchedules))
					So(result.Schedules, ShouldHaveLength, len(expectedSchedules))
					for _, schedule := range result.Schedules {
						So(expectedSchedules, ShouldContain, schedule)
					}
				})
			})
			Convey("When injection fail", func() {
				monkey.Patch(repository.FetchManyByBus, func(busID int64) ([]*model.Schedule, error) {
					return nil, errors.New("fail")
				})
				defer monkey.Unpatch(repository.FetchManyByBus)

				query := queries["by-id"]
				So(query, ShouldNotBeEmpty)

				mock.ExpectQuery(regexp.QuoteMeta(query)).WithArgs(int64(1)).WillReturnRows(
					sqlmock.NewRows([]string{"id", "number", "name", "fare"}).
						AddRow(1, "4988", "Bus One", 3.1))
				result, err := Get(1, true)

				So(err, ShouldBeNil)
				So(mock.ExpectationsWereMet(), ShouldBeNil)
				Convey("Schedules must be injected", func() {
					So(result.Schedules, ShouldHaveLength, 0)
				})
			})
			Convey("When the database return a empty result", func() {
				expectedSchedules := []*model.Schedule{
					new(model.Schedule),
					new(model.Schedule),
					new(model.Schedule),
				}
				monkey.Patch(repository.FetchManyByBus, func(busID int64) ([]*model.Schedule, error) {
					return expectedSchedules, nil
				})
				defer monkey.Unpatch(repository.FetchManyByBus)

				query := queries["by-id"]
				So(query, ShouldNotBeEmpty)

				mock.ExpectQuery(regexp.QuoteMeta(query)).WithArgs(int64(1)).WillReturnRows(
					sqlmock.NewRows([]string{"id", "number", "name", "fare"}))
				result, err := Get(1, true)

				So(err, ShouldBeNil)
				So(result, ShouldBeNil)
				So(mock.ExpectationsWereMet(), ShouldBeNil)
			})
		})
		Convey("When database fail", func() {
			Convey("When the schedules is not required", func() {
				query := queries["by-id"]
				So(query, ShouldNotBeEmpty)

				expected := errors.New("fail")
				mock.ExpectQuery(regexp.QuoteMeta(query)).WithArgs(int64(1)).
					WillReturnError(expected)

				result, err := Get(1, false)
				So(err, ShouldEqual, expected)
				So(result, ShouldBeNil)
			})
		})
		Convey("When database return invalid values", func() {
			Convey("It should fail and return the correct error", func() {
				query := queries["by-id"]
				So(query, ShouldNotBeEmpty)

				mock.ExpectQuery(regexp.QuoteMeta(query)).WithArgs(int64(1)).WillReturnRows(
					sqlmock.NewRows([]string{"id", "number", "name", "fare"}).
						AddRow("a", "a", 1, "a"))
				result, err := Get(1, false)

				So(err, ShouldNotBeNil)
				So(result, ShouldBeNil)
				So(mock.ExpectationsWereMet(), ShouldBeNil)
			})
		})
	})
}
