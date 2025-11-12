package repositories

import (
	"database/sql"
	"errors"
	"fmt"
	"my-app-go/models"
	"regexp"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/require"
)

func TestGetAll(t *testing.T) {
	db, mock, err := sqlmock.New()
	require.NoError(t, err)
	defer db.Close()

	rows := sqlmock.NewRows([]string{"id", "name"}).
		AddRow(0, "Профессор Преображенский").
		AddRow(1, "Доктор Борменталь")

	mock.ExpectQuery("SELECT id, name FROM teachers").
		WillReturnRows(rows)

	r := NewTeacherRepository(db)

	teachers, err := r.GetAll()
	require.NoError(t, err)
	require.Len(t, teachers, 2)
	require.Equal(t, "Профессор Преображенский", teachers[0].Name)
	require.Equal(t, "Доктор Борменталь", teachers[1].Name)

	err = mock.ExpectationsWereMet()
	require.NoError(t, err)
}

func TestInsert(t *testing.T) {
	db, mock, err := sqlmock.New()
	require.NoError(t, err)
	defer db.Close()

	r := NewTeacherRepository(db)

	testTable := []struct {
		name          string
		inputName     string
		wantErr       bool
		mockSetupFunc func()
	}{
		{
			name:      "Teacher inserted successfully",
			inputName: "Профессор Преображенский",
			wantErr:   false,
			mockSetupFunc: func() {
				mock.ExpectExec("INSERT INTO teachers").
					WithArgs("Профессор Преображенский").
					WillReturnResult(sqlmock.NewResult(1, 1))
			},
		},
		{
			name:      "Returns Insert error",
			inputName: "Доктор Борменталь",
			wantErr:   true,
			mockSetupFunc: func() {
				mock.ExpectExec("INSERT INTO teachers").
					WithArgs("Доктор Борменталь").
					WillReturnError(fmt.Errorf("Insert error"))
			},
		},
	}

	for _, testCase := range testTable {
		t.Run(testCase.name, func(t *testing.T) {
			testCase.mockSetupFunc()

			err := r.Insert(testCase.inputName)

			if testCase.wantErr {
				require.Error(t, err)
			} else {
				require.NoError(t, err)
			}

			err = mock.ExpectationsWereMet()
			require.NoError(t, err)
		})
	}
}

func TestGetById(t *testing.T) {
	db, mock, err := sqlmock.New()
	require.NoError(t, err)
	defer db.Close()

	r := NewTeacherRepository(db)

	testTable := []struct {
		name           string
		inputId        string
		expectedResult models.Teacher
		wantErr        bool
		mockSetupFunc  func()
	}{
		{
			name:           "Teacher found by ID",
			inputId:        "0",
			expectedResult: models.Teacher{ID: 0, Name: "Профессор Преображенский"},
			wantErr:        false,
			mockSetupFunc: func() {
				rows := sqlmock.NewRows([]string{"id", "name"}).
					AddRow(0, "Профессор Преображенский")

				mock.ExpectQuery("SELECT id, name FROM teachers WHERE id = ?").
					WithArgs("0").
					WillReturnRows(rows)
			},
		},
		{
			name:           "Error: Teacher not found",
			inputId:        "123",
			expectedResult: models.Teacher{},
			wantErr:        true,
			mockSetupFunc: func() {
				mock.ExpectQuery("SELECT id, name FROM teachers WHERE id = ?").
					WithArgs("123").
					WillReturnError(sql.ErrNoRows)
			},
		},
	}

	for _, testCase := range testTable {
		t.Run(testCase.name, func(t *testing.T) {
			testCase.mockSetupFunc()

			teacher, err := r.GetById(testCase.inputId)

			if testCase.wantErr {
				require.Error(t, err)
			} else {
				require.NoError(t, err)
				require.Equal(t, testCase.expectedResult, teacher)
			}

			err = mock.ExpectationsWereMet()
			require.NoError(t, err)
		})
	}
}

func TestUpdate(t *testing.T) {
	db, mock, err := sqlmock.New()
	require.NoError(t, err)
	defer db.Close()

	r := NewTeacherRepository(db)

	type args struct {
		id   string
		name string
	}
	testTable := []struct {
		name          string
		args          args
		wantErr       bool
		mockSetupFunc func()
	}{
		{
			name: "Successful update",
			args: args{
				id:   "0",
				name: "Доктор Айболит",
			},
			wantErr: false,
			mockSetupFunc: func() {
				sqlmock.NewRows([]string{"id", "name"}).
					AddRow(0, "Профессор Преображенский")

				mock.ExpectExec(regexp.QuoteMeta(
					"UPDATE teachers SET name = ? WHERE id = ?")).
					WithArgs("Доктор Айболит", "0").
					WillReturnResult(sqlmock.NewResult(0, 1))
			},
		},
		{
			name: "Database update failed",
			args: args{
				id:   "123",
				name: "Доктор Айболит",
			},
			wantErr: true,
			mockSetupFunc: func() {
				sqlmock.NewRows([]string{"id", "name"}).
					AddRow(0, "Профессор Преображенский")

				mock.ExpectExec(regexp.QuoteMeta(
					"UPDATE teachers SET name = ? WHERE id = ?")).
					WithArgs("Доктор Айболит", "123").
					WillReturnError(errors.New("Database update failed"))
			},
		},
	}

	for _, testCase := range testTable {
		t.Run(testCase.name, func(t *testing.T) {
			testCase.mockSetupFunc()

			err := r.Update(testCase.args.id, testCase.args.name)

			if testCase.wantErr {
				require.Error(t, err)
			} else {
				require.NoError(t, err)
			}

			err = mock.ExpectationsWereMet()
			require.NoError(t, err)
		})
	}
}
