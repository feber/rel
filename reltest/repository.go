package reltest

import (
	"runtime"
	"testing"

	"github.com/Fs02/rel"
	"github.com/stretchr/testify/mock"
)

func must(err error) {
	if err != nil {
		panic(err)
	}
}

// Repository is an autogenerated mock type for the Repository type
type Repository struct {
	mock mock.Mock
	tx   *Repository
}

var _ rel.Repository = (*Repository)(nil)

// Adapter provides a mock function with given fields:
func (r *Repository) Adapter() rel.Adapter {
	return nil
}

// SetLogger provides a mock function with given fields: logger
func (r *Repository) SetLogger(logger ...rel.Logger) {
}

// Aggregate provides a mock function with given fields: query, aggregate, field
func (r *Repository) Aggregate(query rel.Query, aggregate string, field string) (int, error) {
	ret := r.mock.Called(query, aggregate, field)
	return ret.Int(0), ret.Error(1)
}

// MustAggregate provides a mock function with given fields: query, aggregate, field
func (r *Repository) MustAggregate(query rel.Query, aggregate string, field string) int {
	result, err := r.Aggregate(query, aggregate, field)
	must(err)
	return result
}

// ExpectAggregate apply mocks and expectations for Aggregate
func (r *Repository) ExpectAggregate(query rel.Query, aggregate string, field string) *Aggregate {
	return ExpectAggregate(r, query, aggregate, field)
}

// Count provides a mock function with given fields: collection, queriers
func (r *Repository) Count(collection string, queriers ...rel.Querier) (int, error) {
	ret := r.mock.Called(collection, queriers)
	return ret.Int(0), ret.Error(1)
}

// MustCount provides a mock function with given fields: collection, queriers
func (r *Repository) MustCount(collection string, queriers ...rel.Querier) int {
	count, err := r.Count(collection, queriers...)
	must(err)
	return count
}

// ExpectCount apply mocks and expectations for Count
func (r *Repository) ExpectCount(collection string, queriers ...rel.Querier) *Aggregate {
	return ExpectCount(r, collection, queriers)
}

// Find provides a mock function with given fields: record, queriers
func (r *Repository) Find(record interface{}, queriers ...rel.Querier) error {
	return r.mock.Called(record, queriers).Error(0)
}

// MustFind provides a mock function with given fields: record, queriers
func (r *Repository) MustFind(record interface{}, queriers ...rel.Querier) {
	must(r.Find(record, queriers...))
}

// ExpectFind apply mocks and expectations for Find
func (r *Repository) ExpectFind(queriers ...rel.Querier) *Find {
	return ExpectFind(r, queriers)
}

// FindAll provides a mock function with given fields: records, queriers
func (r *Repository) FindAll(records interface{}, queriers ...rel.Querier) error {
	return r.mock.Called(records, queriers).Error(0)
}

// ExpectFindAll apply mocks and expectations for FindAll
func (r *Repository) ExpectFindAll(queriers ...rel.Querier) *FindAll {
	return ExpectFindAll(r, queriers)
}

// MustFindAll provides a mock function with given fields: records, queriers
func (r *Repository) MustFindAll(records interface{}, queriers ...rel.Querier) {
	must(r.FindAll(records, queriers...))
}

// Insert provides a mock function with given fields: record, changers
func (r *Repository) Insert(record interface{}, changers ...rel.Changer) error {
	return r.mock.Called(record, changers).Error(0)
}

// MustInsert provides a mock function with given fields: record, changers
func (r *Repository) MustInsert(record interface{}, changers ...rel.Changer) {
	must(r.Insert(record, changers...))
}

// ExpectInsert apply mocks and expectations for Insert
func (r *Repository) ExpectInsert(changers ...rel.Changer) *Modify {
	return ExpectModify(r, "Insert", changers, true)
}

// InsertAll provides a mock function with given fields: records, changes
func (r *Repository) InsertAll(records interface{}, changes ...rel.Changes) error {
	return r.mock.Called(records, changes).Error(0)
}

// MustInsertAll provides a mock function with given fields: records, changes
func (r *Repository) MustInsertAll(records interface{}, changes ...rel.Changes) {
	must(r.InsertAll(records, changes...))
}

// ExpectInsertAll apply mocks and expectations for InsertAll
func (r *Repository) ExpectInsertAll(changes ...rel.Changes) *Modify {
	return ExpectInsertAll(r, changes)
}

// Update provides a mock function with given fields: record, changers
func (r *Repository) Update(record interface{}, changers ...rel.Changer) error {
	return r.mock.Called(record, changers).Error(0)
}

// MustUpdate provides a mock function with given fields: record, changers
func (r *Repository) MustUpdate(record interface{}, changers ...rel.Changer) {
	must(r.Update(record, changers...))
}

// ExpectUpdate apply mocks and expectations for Update
func (r *Repository) ExpectUpdate(changers ...rel.Changer) *Modify {
	return ExpectModify(r, "Update", changers, false)
}

// Delete provides a mock function with given fields: record
func (r *Repository) Delete(record interface{}) error {
	return r.mock.Called(record).Error(0)
}

// MustDelete provides a mock function with given fields: record
func (r *Repository) MustDelete(record interface{}) {
	must(r.Delete(record))
}

// ExpectDelete apply mocks and expectations for Delete
func (r *Repository) ExpectDelete() *Delete {
	return ExpectDelete(r)
}

// DeleteAll provides a mock function with given fields: queriers
func (r *Repository) DeleteAll(queriers ...rel.Querier) error {
	return r.mock.Called(queriers).Error(0)
}

// MustDeleteAll provides a mock function with given fields: queriers
func (r *Repository) MustDeleteAll(queriers ...rel.Querier) {
	must(r.DeleteAll(queriers...))
}

// ExpectDeleteAll apply mocks and expectations for DeleteAll
func (r *Repository) ExpectDeleteAll(queriers ...rel.Querier) *DeleteAll {
	return ExpectDeleteAll(r, queriers)
}

// Preload provides a mock function with given fields: records, field, queriers
func (r *Repository) Preload(records interface{}, field string, queriers ...rel.Querier) error {
	return r.mock.Called(records, field, queriers).Error(0)
}

// MustPreload provides a mock function with given fields: records, field, queriers
func (r *Repository) MustPreload(records interface{}, field string, queriers ...rel.Querier) {
	must(r.Preload(records, field, queriers...))
}

// ExpectPreload apply mocks and expectations for Preload
func (r *Repository) ExpectPreload(field string, queriers ...rel.Querier) *Preload {
	return ExpectPreload(r, field, queriers)
}

// Transaction provides a mock function with given fields: fn
func (r *Repository) Transaction(fn func(rel.Repository) error) error {
	r.mock.Called()

	var err error
	func() {
		defer func() {
			if p := recover(); p != nil {
				switch e := p.(type) {
				case runtime.Error:
					panic(e)
				case error:
					err = e
				default:
					panic(e)
				}
			}
		}()

		err = fn(r.tx)
	}()

	return err
}

// ExpectTransaction declare expectation inside transaction.
func (r *Repository) ExpectTransaction(fn func(*Repository)) {
	r.mock.On("Transaction").Once()

	if r.tx == nil {
		r.tx = &Repository{}
	}

	fn(r.tx)
}

// AssertExpectations asserts that everything was in fact called as expected. Calls may have occurred in any order.
func (r *Repository) AssertExpectations(t *testing.T) bool {
	if r.tx != nil {
		return r.mock.AssertExpectations(t) && r.tx.AssertExpectations(t)
	}

	return r.mock.AssertExpectations(t)
}