// Code generated by gen_tests.py and process_polyglot.py.
// Do not edit this file directly.
// The template for this file is located at:
// ../template.go.tpl
package reql_tests

import (
	"testing"
	"time"

	"github.com/stretchr/testify/suite"
	r "gopkg.in/gorethink/gorethink.v3"
	"gopkg.in/gorethink/gorethink.v3/internal/compare"
)

// Test geo constructors
func TestGeoConstructorsSuite(t *testing.T) {
	suite.Run(t, new(GeoConstructorsSuite))
}

type GeoConstructorsSuite struct {
	suite.Suite

	session *r.Session
}

func (suite *GeoConstructorsSuite) SetupTest() {
	suite.T().Log("Setting up GeoConstructorsSuite")
	// Use imports to prevent errors
	_ = time.Time{}
	_ = compare.AnythingIsFine

	session, err := r.Connect(r.ConnectOpts{
		Address: url,
	})
	suite.Require().NoError(err, "Error returned when connecting to server")
	suite.session = session

	r.DBDrop("test").Exec(suite.session)
	err = r.DBCreate("test").Exec(suite.session)
	suite.Require().NoError(err)
	err = r.DB("test").Wait().Exec(suite.session)
	suite.Require().NoError(err)

}

func (suite *GeoConstructorsSuite) TearDownSuite() {
	suite.T().Log("Tearing down GeoConstructorsSuite")

	if suite.session != nil {
		r.DB("rethinkdb").Table("_debug_scratch").Delete().Exec(suite.session)
		r.DBDrop("test").Exec(suite.session)

		suite.session.Close()
	}
}

func (suite *GeoConstructorsSuite) TestCases() {
	suite.T().Log("Running GeoConstructorsSuite: Test geo constructors")

	{
		// geo/constructors.yaml line #4
		/* ({'$reql_type$':'GEOMETRY', 'coordinates':[0, 0], 'type':'Point'}) */
		var expected_ map[interface{}]interface{} = map[interface{}]interface{}{"$reql_type$": "GEOMETRY", "coordinates": []interface{}{0, 0}, "type": "Point"}
		/* r.point(0, 0) */

		suite.T().Log("About to run line #4: r.Point(0, 0)")

		runAndAssert(suite.Suite, expected_, r.Point(0, 0), suite.session, r.RunOpts{
			GeometryFormat: "raw",
			GroupFormat:    "map",
		})
		suite.T().Log("Finished running line #4")
	}

	{
		// geo/constructors.yaml line #6
		/* ({'$reql_type$':'GEOMETRY', 'coordinates':[0, -90], 'type':'Point'}) */
		var expected_ map[interface{}]interface{} = map[interface{}]interface{}{"$reql_type$": "GEOMETRY", "coordinates": []interface{}{0, -90}, "type": "Point"}
		/* r.point(0, -90) */

		suite.T().Log("About to run line #6: r.Point(0, -90)")

		runAndAssert(suite.Suite, expected_, r.Point(0, -90), suite.session, r.RunOpts{
			GeometryFormat: "raw",
			GroupFormat:    "map",
		})
		suite.T().Log("Finished running line #6")
	}

	{
		// geo/constructors.yaml line #8
		/* ({'$reql_type$':'GEOMETRY', 'coordinates':[0, 90], 'type':'Point'}) */
		var expected_ map[interface{}]interface{} = map[interface{}]interface{}{"$reql_type$": "GEOMETRY", "coordinates": []interface{}{0, 90}, "type": "Point"}
		/* r.point(0, 90) */

		suite.T().Log("About to run line #8: r.Point(0, 90)")

		runAndAssert(suite.Suite, expected_, r.Point(0, 90), suite.session, r.RunOpts{
			GeometryFormat: "raw",
			GroupFormat:    "map",
		})
		suite.T().Log("Finished running line #8")
	}

	{
		// geo/constructors.yaml line #10
		/* ({'$reql_type$':'GEOMETRY', 'coordinates':[-180, 0], 'type':'Point'}) */
		var expected_ map[interface{}]interface{} = map[interface{}]interface{}{"$reql_type$": "GEOMETRY", "coordinates": []interface{}{-180, 0}, "type": "Point"}
		/* r.point(-180, 0) */

		suite.T().Log("About to run line #10: r.Point(-180, 0)")

		runAndAssert(suite.Suite, expected_, r.Point(-180, 0), suite.session, r.RunOpts{
			GeometryFormat: "raw",
			GroupFormat:    "map",
		})
		suite.T().Log("Finished running line #10")
	}

	{
		// geo/constructors.yaml line #12
		/* ({'$reql_type$':'GEOMETRY', 'coordinates':[180, 0], 'type':'Point'}) */
		var expected_ map[interface{}]interface{} = map[interface{}]interface{}{"$reql_type$": "GEOMETRY", "coordinates": []interface{}{180, 0}, "type": "Point"}
		/* r.point(180, 0) */

		suite.T().Log("About to run line #12: r.Point(180, 0)")

		runAndAssert(suite.Suite, expected_, r.Point(180, 0), suite.session, r.RunOpts{
			GeometryFormat: "raw",
			GroupFormat:    "map",
		})
		suite.T().Log("Finished running line #12")
	}

	{
		// geo/constructors.yaml line #14
		/* err('ReqlQueryLogicError', 'Latitude must be between -90 and 90.  Got -91.', [0]) */
		var expected_ Err = err("ReqlQueryLogicError", "Latitude must be between -90 and 90.  Got -91.")
		/* r.point(0, -91) */

		suite.T().Log("About to run line #14: r.Point(0, -91)")

		runAndAssert(suite.Suite, expected_, r.Point(0, -91), suite.session, r.RunOpts{
			GeometryFormat: "raw",
			GroupFormat:    "map",
		})
		suite.T().Log("Finished running line #14")
	}

	{
		// geo/constructors.yaml line #16
		/* err('ReqlQueryLogicError', 'Latitude must be between -90 and 90.  Got 91.', [0]) */
		var expected_ Err = err("ReqlQueryLogicError", "Latitude must be between -90 and 90.  Got 91.")
		/* r.point(0, 91) */

		suite.T().Log("About to run line #16: r.Point(0, 91)")

		runAndAssert(suite.Suite, expected_, r.Point(0, 91), suite.session, r.RunOpts{
			GeometryFormat: "raw",
			GroupFormat:    "map",
		})
		suite.T().Log("Finished running line #16")
	}

	{
		// geo/constructors.yaml line #18
		/* err('ReqlQueryLogicError', 'Longitude must be between -180 and 180.  Got -181.', [0]) */
		var expected_ Err = err("ReqlQueryLogicError", "Longitude must be between -180 and 180.  Got -181.")
		/* r.point(-181, 0) */

		suite.T().Log("About to run line #18: r.Point(-181, 0)")

		runAndAssert(suite.Suite, expected_, r.Point(-181, 0), suite.session, r.RunOpts{
			GeometryFormat: "raw",
			GroupFormat:    "map",
		})
		suite.T().Log("Finished running line #18")
	}

	{
		// geo/constructors.yaml line #20
		/* err('ReqlQueryLogicError', 'Longitude must be between -180 and 180.  Got 181.', [0]) */
		var expected_ Err = err("ReqlQueryLogicError", "Longitude must be between -180 and 180.  Got 181.")
		/* r.point(181, 0) */

		suite.T().Log("About to run line #20: r.Point(181, 0)")

		runAndAssert(suite.Suite, expected_, r.Point(181, 0), suite.session, r.RunOpts{
			GeometryFormat: "raw",
			GroupFormat:    "map",
		})
		suite.T().Log("Finished running line #20")
	}

	{
		// geo/constructors.yaml line #28
		/* err('ReqlQueryLogicError', 'Invalid LineString.  Are there antipodal or duplicate vertices?', [0]) */
		var expected_ Err = err("ReqlQueryLogicError", "Invalid LineString.  Are there antipodal or duplicate vertices?")
		/* r.line([0,0], [0,0]) */

		suite.T().Log("About to run line #28: r.Line([]interface{}{0, 0}, []interface{}{0, 0})")

		runAndAssert(suite.Suite, expected_, r.Line([]interface{}{0, 0}, []interface{}{0, 0}), suite.session, r.RunOpts{
			GeometryFormat: "raw",
			GroupFormat:    "map",
		})
		suite.T().Log("Finished running line #28")
	}

	{
		// geo/constructors.yaml line #30
		/* ({'$reql_type$':'GEOMETRY', 'coordinates':[[0,0], [0,1]], 'type':'LineString'}) */
		var expected_ map[interface{}]interface{} = map[interface{}]interface{}{"$reql_type$": "GEOMETRY", "coordinates": []interface{}{[]interface{}{0, 0}, []interface{}{0, 1}}, "type": "LineString"}
		/* r.line([0,0], [0,1]) */

		suite.T().Log("About to run line #30: r.Line([]interface{}{0, 0}, []interface{}{0, 1})")

		runAndAssert(suite.Suite, expected_, r.Line([]interface{}{0, 0}, []interface{}{0, 1}), suite.session, r.RunOpts{
			GeometryFormat: "raw",
			GroupFormat:    "map",
		})
		suite.T().Log("Finished running line #30")
	}

	{
		// geo/constructors.yaml line #32
		/* err('ReqlQueryLogicError', 'Expected point coordinate pair.  Got 1 element array instead of a 2 element one.', [0]) */
		var expected_ Err = err("ReqlQueryLogicError", "Expected point coordinate pair.  Got 1 element array instead of a 2 element one.")
		/* r.line([0,0], [1]) */

		suite.T().Log("About to run line #32: r.Line([]interface{}{0, 0}, []interface{}{1})")

		runAndAssert(suite.Suite, expected_, r.Line([]interface{}{0, 0}, []interface{}{1}), suite.session, r.RunOpts{
			GeometryFormat: "raw",
			GroupFormat:    "map",
		})
		suite.T().Log("Finished running line #32")
	}

	{
		// geo/constructors.yaml line #34
		/* err('ReqlQueryLogicError', 'Expected point coordinate pair.  Got 3 element array instead of a 2 element one.', [0]) */
		var expected_ Err = err("ReqlQueryLogicError", "Expected point coordinate pair.  Got 3 element array instead of a 2 element one.")
		/* r.line([0,0], [1,0,0]) */

		suite.T().Log("About to run line #34: r.Line([]interface{}{0, 0}, []interface{}{1, 0, 0})")

		runAndAssert(suite.Suite, expected_, r.Line([]interface{}{0, 0}, []interface{}{1, 0, 0}), suite.session, r.RunOpts{
			GeometryFormat: "raw",
			GroupFormat:    "map",
		})
		suite.T().Log("Finished running line #34")
	}

	{
		// geo/constructors.yaml line #36
		/* ({'$reql_type$':'GEOMETRY', 'coordinates':[[0,0], [0,1], [0,0]], 'type':'LineString'}) */
		var expected_ map[interface{}]interface{} = map[interface{}]interface{}{"$reql_type$": "GEOMETRY", "coordinates": []interface{}{[]interface{}{0, 0}, []interface{}{0, 1}, []interface{}{0, 0}}, "type": "LineString"}
		/* r.line([0,0], [0,1], [0,0]) */

		suite.T().Log("About to run line #36: r.Line([]interface{}{0, 0}, []interface{}{0, 1}, []interface{}{0, 0})")

		runAndAssert(suite.Suite, expected_, r.Line([]interface{}{0, 0}, []interface{}{0, 1}, []interface{}{0, 0}), suite.session, r.RunOpts{
			GeometryFormat: "raw",
			GroupFormat:    "map",
		})
		suite.T().Log("Finished running line #36")
	}

	{
		// geo/constructors.yaml line #38
		/* ({'$reql_type$':'GEOMETRY', 'coordinates':[[0,0], [0,1], [0,0]], 'type':'LineString'}) */
		var expected_ map[interface{}]interface{} = map[interface{}]interface{}{"$reql_type$": "GEOMETRY", "coordinates": []interface{}{[]interface{}{0, 0}, []interface{}{0, 1}, []interface{}{0, 0}}, "type": "LineString"}
		/* r.line(r.point(0,0), r.point(0,1), r.point(0,0)) */

		suite.T().Log("About to run line #38: r.Line(r.Point(0, 0), r.Point(0, 1), r.Point(0, 0))")

		runAndAssert(suite.Suite, expected_, r.Line(r.Point(0, 0), r.Point(0, 1), r.Point(0, 0)), suite.session, r.RunOpts{
			GeometryFormat: "raw",
			GroupFormat:    "map",
		})
		suite.T().Log("Finished running line #38")
	}

	{
		// geo/constructors.yaml line #40
		/* err('ReqlQueryLogicError', 'Expected geometry of type `Point` but found `LineString`.', [0]) */
		var expected_ Err = err("ReqlQueryLogicError", "Expected geometry of type `Point` but found `LineString`.")
		/* r.line(r.point(0,0), r.point(1,0), r.line([0,0], [1,0])) */

		suite.T().Log("About to run line #40: r.Line(r.Point(0, 0), r.Point(1, 0), r.Line([]interface{}{0, 0}, []interface{}{1, 0}))")

		runAndAssert(suite.Suite, expected_, r.Line(r.Point(0, 0), r.Point(1, 0), r.Line([]interface{}{0, 0}, []interface{}{1, 0})), suite.session, r.RunOpts{
			GeometryFormat: "raw",
			GroupFormat:    "map",
		})
		suite.T().Log("Finished running line #40")
	}

	{
		// geo/constructors.yaml line #50
		/* err('ReqlQueryLogicError', 'Invalid LinearRing.  Are there antipodal or duplicate vertices? Is it self-intersecting?', [0]) */
		var expected_ Err = err("ReqlQueryLogicError", "Invalid LinearRing.  Are there antipodal or duplicate vertices? Is it self-intersecting?")
		/* r.polygon([0,0], [0,0], [0,0], [0,0]) */

		suite.T().Log("About to run line #50: r.Polygon([]interface{}{0, 0}, []interface{}{0, 0}, []interface{}{0, 0}, []interface{}{0, 0})")

		runAndAssert(suite.Suite, expected_, r.Polygon([]interface{}{0, 0}, []interface{}{0, 0}, []interface{}{0, 0}, []interface{}{0, 0}), suite.session, r.RunOpts{
			GeometryFormat: "raw",
			GroupFormat:    "map",
		})
		suite.T().Log("Finished running line #50")
	}

	{
		// geo/constructors.yaml line #52
		/* ({'$reql_type$':'GEOMETRY', 'coordinates':[[[0,0], [0,1], [1,0], [0,0]]], 'type':'Polygon'}) */
		var expected_ map[interface{}]interface{} = map[interface{}]interface{}{"$reql_type$": "GEOMETRY", "coordinates": []interface{}{[]interface{}{[]interface{}{0, 0}, []interface{}{0, 1}, []interface{}{1, 0}, []interface{}{0, 0}}}, "type": "Polygon"}
		/* r.polygon([0,0], [0,1], [1,0]) */

		suite.T().Log("About to run line #52: r.Polygon([]interface{}{0, 0}, []interface{}{0, 1}, []interface{}{1, 0})")

		runAndAssert(suite.Suite, expected_, r.Polygon([]interface{}{0, 0}, []interface{}{0, 1}, []interface{}{1, 0}), suite.session, r.RunOpts{
			GeometryFormat: "raw",
			GroupFormat:    "map",
		})
		suite.T().Log("Finished running line #52")
	}

	{
		// geo/constructors.yaml line #54
		/* ({'$reql_type$':'GEOMETRY', 'coordinates':[[[0,0], [0,1], [1,0], [0,0]]], 'type':'Polygon'}) */
		var expected_ map[interface{}]interface{} = map[interface{}]interface{}{"$reql_type$": "GEOMETRY", "coordinates": []interface{}{[]interface{}{[]interface{}{0, 0}, []interface{}{0, 1}, []interface{}{1, 0}, []interface{}{0, 0}}}, "type": "Polygon"}
		/* r.polygon([0,0], [0,1], [1,0], [0,0]) */

		suite.T().Log("About to run line #54: r.Polygon([]interface{}{0, 0}, []interface{}{0, 1}, []interface{}{1, 0}, []interface{}{0, 0})")

		runAndAssert(suite.Suite, expected_, r.Polygon([]interface{}{0, 0}, []interface{}{0, 1}, []interface{}{1, 0}, []interface{}{0, 0}), suite.session, r.RunOpts{
			GeometryFormat: "raw",
			GroupFormat:    "map",
		})
		suite.T().Log("Finished running line #54")
	}

	{
		// geo/constructors.yaml line #56
		/* err('ReqlQueryLogicError', 'Invalid LinearRing.  Are there antipodal or duplicate vertices? Is it self-intersecting?', [0]) */
		var expected_ Err = err("ReqlQueryLogicError", "Invalid LinearRing.  Are there antipodal or duplicate vertices? Is it self-intersecting?")
		/* r.polygon([0,0], [0,1], [1,0], [-1,0.5]) */

		suite.T().Log("About to run line #56: r.Polygon([]interface{}{0, 0}, []interface{}{0, 1}, []interface{}{1, 0}, []interface{}{-1, 0.5})")

		runAndAssert(suite.Suite, expected_, r.Polygon([]interface{}{0, 0}, []interface{}{0, 1}, []interface{}{1, 0}, []interface{}{-1, 0.5}), suite.session, r.RunOpts{
			GeometryFormat: "raw",
			GroupFormat:    "map",
		})
		suite.T().Log("Finished running line #56")
	}

	{
		// geo/constructors.yaml line #58
		/* err('ReqlQueryLogicError', 'Expected point coordinate pair.  Got 1 element array instead of a 2 element one.', [0]) */
		var expected_ Err = err("ReqlQueryLogicError", "Expected point coordinate pair.  Got 1 element array instead of a 2 element one.")
		/* r.polygon([0,0], [0,1], [0]) */

		suite.T().Log("About to run line #58: r.Polygon([]interface{}{0, 0}, []interface{}{0, 1}, []interface{}{0})")

		runAndAssert(suite.Suite, expected_, r.Polygon([]interface{}{0, 0}, []interface{}{0, 1}, []interface{}{0}), suite.session, r.RunOpts{
			GeometryFormat: "raw",
			GroupFormat:    "map",
		})
		suite.T().Log("Finished running line #58")
	}

	{
		// geo/constructors.yaml line #60
		/* err('ReqlQueryLogicError', 'Expected point coordinate pair.  Got 3 element array instead of a 2 element one.', [0]) */
		var expected_ Err = err("ReqlQueryLogicError", "Expected point coordinate pair.  Got 3 element array instead of a 2 element one.")
		/* r.polygon([0,0], [0,1], [0,1,0]) */

		suite.T().Log("About to run line #60: r.Polygon([]interface{}{0, 0}, []interface{}{0, 1}, []interface{}{0, 1, 0})")

		runAndAssert(suite.Suite, expected_, r.Polygon([]interface{}{0, 0}, []interface{}{0, 1}, []interface{}{0, 1, 0}), suite.session, r.RunOpts{
			GeometryFormat: "raw",
			GroupFormat:    "map",
		})
		suite.T().Log("Finished running line #60")
	}

	{
		// geo/constructors.yaml line #62
		/* err('ReqlQueryLogicError', 'Expected geometry of type `Point` but found `LineString`.', [0]) */
		var expected_ Err = err("ReqlQueryLogicError", "Expected geometry of type `Point` but found `LineString`.")
		/* r.polygon(r.point(0,0), r.point(0,1), r.line([0,0], [0,1])) */

		suite.T().Log("About to run line #62: r.Polygon(r.Point(0, 0), r.Point(0, 1), r.Line([]interface{}{0, 0}, []interface{}{0, 1}))")

		runAndAssert(suite.Suite, expected_, r.Polygon(r.Point(0, 0), r.Point(0, 1), r.Line([]interface{}{0, 0}, []interface{}{0, 1})), suite.session, r.RunOpts{
			GeometryFormat: "raw",
			GroupFormat:    "map",
		})
		suite.T().Log("Finished running line #62")
	}
}
