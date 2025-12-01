package patient

import (
	"context"
	"regexp"
	"testing"
	"time"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/require"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	"github.com/thanamin/go-api/internal/infrastructures/dataMapper"
	repoPatient "github.com/thanamin/go-api/internal/infrastructures/persistence/repositories/patient"
)

// helper to open gorm DB with sqlmock
func openMockDB(t *testing.T) (*gorm.DB, sqlmock.Sqlmock, func()) {
	db, mock, err := sqlmock.New()
	require.NoError(t, err)

	dialector := postgres.New(postgres.Config{
		Conn: db,
	})

	gdb, err := gorm.Open(dialector, &gorm.Config{})
	require.NoError(t, err)

	closeFn := func() {
		_ = db.Close()
	}
	return gdb, mock, closeFn
}

func TestExecute_ReturnsPatients(t *testing.T) {
	gdb, mock, closeFn := openMockDB(t)
	defer closeFn()

	// expect a SELECT with WHERE and LIMIT/OFFSET; allow flexible WHERE via regexp
	rows := sqlmock.NewRows([]string{"id", "hospital_id", "national_id", "first_name_th", "last_name_th", "date_of_birth", "patient_hn", "phone_number", "email", "gender", "created_at", "updated_at"}).AddRow(
		1, 1, "1234567890123", "Somchai", "Suksai", time.Now(), "HN001", "0812345678", "a@b.c", "M", time.Now(), time.Now(),
	)

	// build expected query regex — match deleted_at condition and LIMIT (OFFSET is omitted when 0)
	mock.ExpectQuery(regexp.QuoteMeta("SELECT * FROM \"patients\" WHERE deleted_at IS NULL LIMIT $1")).WillReturnRows(rows)

	mapper := dataMapper.NewPatientMapper()
	// create concrete repo that implements DBWithContext
	patientRepo := repoPatient.NewPatientRepository(gdb)
	uc := InitSearch(patientRepo, mapper)

	ctx := context.Background()
	params := &SearchParams{Limit: 1, Offset: 0}

	patients, err := uc.Execute(ctx, params)
	require.NoError(t, err)
	require.Len(t, patients, 1)
	require.Equal(t, "1234567890123", patients[0].NationalID)

	require.NoError(t, mock.ExpectationsWereMet())
}

func TestExecute_InvalidDate_ReturnsError(t *testing.T) {
	gdb, _, closeFn := openMockDB(t)
	defer closeFn()

	mapper := dataMapper.NewPatientMapper()
	patientRepo := repoPatient.NewPatientRepository(gdb)
	uc := InitSearch(patientRepo, mapper)

	ctx := context.Background()
	params := &SearchParams{DateOfBirth: "invalid-date"}

	_, err := uc.Execute(ctx, params)
	require.Error(t, err)
}
