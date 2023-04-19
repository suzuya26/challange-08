package service

import (
	"errors"
	"sesi_8/model"
	"sesi_8/repository/mocks"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

func Test_BookService_GetBookById(t *testing.T) {
	type testCase struct {
		name           string
		wantError      bool
		expectedResult model.Book
		expectedError  error
		onBookRepo     func(mock *mocks.MockBookRepo)
	}

	var testTable []testCase

	testTable = append(testTable, testCase{
		name:      "success",
		wantError: false,
		onBookRepo: func(mock *mocks.MockBookRepo) {
			mock.EXPECT().GetBookById(gomock.Any()).Return(model.Book{
				ID:          1,
				Title:       "Bahasa Inggris kelas 4",
				Author:      "Bu cut suti",
				Description: "Buku pelajaran bahasa inggris untuk anak kelas 4",
			}, nil).Times(1)
		},
		expectedResult: model.Book{
			ID:          1,
			Title:       "Bahasa Inggris kelas 4",
			Author:      "Bu cut suti",
			Description: "Buku pelajaran bahasa inggris untuk anak kelas 4",
		},
	})

	testTable = append(testTable, testCase{
		name:          "record not found",
		wantError:     true,
		expectedError: errors.New("record not found"),
		onBookRepo: func(mock *mocks.MockBookRepo) {
			mock.EXPECT().GetBookById(gomock.Any()).Return(model.Book{}, errors.New("record not found")).Times(1)
		},
	})

	for _, testCase := range testTable {
		t.Run(testCase.name, func(t *testing.T) {
			mockCtrl := gomock.NewController(t)

			bookRepo := mocks.NewMockBookRepo(mockCtrl)

			if testCase.onBookRepo != nil {
				testCase.onBookRepo(bookRepo)
			}

			service := Service{
				repo: bookRepo,
			}

			res, err := service.GetBookById(1)

			if testCase.wantError {
				assert.EqualError(t, err, testCase.expectedError.Error())
			} else {
				assert.Nil(t, err)
				assert.Equal(t, testCase.expectedResult, res)
			}
		})
	}
}

func Test_BookService_GetAllBook(t *testing.T) {
	type testCase struct {
		name           string
		wantError      bool
		expectedResult []model.Book
		expectedError  error
		onBookRepo     func(mock *mocks.MockBookRepo)
	}

	var testTable []testCase

	testTable = append(testTable, testCase{
		name:      "success",
		wantError: false,
		onBookRepo: func(mock *mocks.MockBookRepo) {
			mock.EXPECT().GetAllBook().Return([]model.Book{
				{
					ID:          1,
					Title:       "Bahasa Inggris kelas 4",
					Author:      "Bu cut suti",
					Description: "Buku pelajaran bahasa inggris untuk anak kelas 4",
				}, {
					ID:          2,
					Title:       "Kamus Bahasa Arab",
					Author:      "Abdoelle",
					Description: "Kamus Bahasa Arab Gaul",
				},
			}, nil).Times(1)
		},
		expectedResult: []model.Book{
			{
				ID:          1,
				Title:       "Bahasa Inggris kelas 4",
				Author:      "Bu cut suti",
				Description: "Buku pelajaran bahasa inggris untuk anak kelas 4",
			}, {
				ID:          2,
				Title:       "Kamus Bahasa Arab",
				Author:      "Abdoelle",
				Description: "Kamus Bahasa Arab Gaul",
			},
		},
	})

	testTable = append(testTable, testCase{
		name:          "record not found",
		wantError:     true,
		expectedError: errors.New("record not found"),
		onBookRepo: func(mock *mocks.MockBookRepo) {
			mock.EXPECT().GetAllBook().Return([]model.Book{}, errors.New("record not found")).Times(1)
		},
	})

	for _, testCase := range testTable {
		t.Run(testCase.name, func(t *testing.T) {
			mockCtrl := gomock.NewController(t)

			bookRepo := mocks.NewMockBookRepo(mockCtrl)

			if testCase.onBookRepo != nil {
				testCase.onBookRepo(bookRepo)
			}

			service := Service{
				repo: bookRepo,
			}

			res, err := service.GetAllBook()

			if testCase.wantError {
				assert.EqualError(t, err, testCase.expectedError.Error())
			} else {
				assert.Nil(t, err)
				assert.Equal(t, testCase.expectedResult, res)
			}
		})
	}

}
