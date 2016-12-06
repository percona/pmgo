package pmgo

import mgo "gopkg.in/mgo.v2"

type QueryManager interface {
	All(result interface{}) error
	//	Apply(change mgo.Change, result interface{}) (info *mgo.ChangeInfo, err error)
	//	Batch(n int) *QueryManager
	//	Comment(comment string) *QueryManager
	Count() (n int, err error)
	//	Distinct(key string, result interface{}) error
	//	Explain(result interface{}) error
	//	For(result interface{}, f func() error) error
	//  Hint(indexKey ...string) *Query
	Iter() *mgo.Iter
	Limit(n int) QueryManager
	//	LogReplay() *QueryManager
	//  MapReduce(job *MapReduce, result interface{}) (info *MapReduceInfo, err error)
	One(result interface{}) (err error)
	//	Prefetch(p float64) *QueryManager
	//	Select(selector interface{}) *QueryManager
	//	SetMaxScan(n int) *QueryManager
	//	SetMaxTime(d time.Duration) *QueryManager
	//	Skip(n int) *QueryManager
	//	Snapshot() *QueryManager
	Sort(fields ...string) QueryManager
	//   Tail(timeout time.Duration) *Iter
}
type Query struct {
	query *mgo.Query
}

func (q *Query) All(result interface{}) error {
	return q.query.All(result)
}

func (q *Query) Count() (int, error) {
	return q.query.Count()
}

func (q *Query) Iter() *mgo.Iter {
	return q.query.Iter()
}

func (q *Query) Limit(n int) QueryManager {
	return &Query{
		query: q.query.Limit(n),
	}
}

func (q *Query) One(result interface{}) error {
	return q.query.One(result)
}

func (q *Query) Sort(fields ...string) QueryManager {
	return &Query{
		query: q.query.Sort(fields...),
	}
}
