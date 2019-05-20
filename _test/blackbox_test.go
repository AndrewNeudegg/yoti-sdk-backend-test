package blackbox

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// TestSanity ensures that the later tests are built on working assumptions.
// If this tests fails the problem is due to the testing paradigm.
// Check to see if the application is running, or if something is using the target port.
func TestSanity(t *testing.T) {
	testHeader(t)
	shim, err := GetShim()
	HandleError(t, err)
	shim.StartServer()
	shim.StopServer()
	testFooter(t)
}

// TestConnection will ensure that the server is accessible.
// If this tests fails the problem is due to the testing paradigm.
// Check to see if the application is running, or if something is using the target port.
func TestConnection(t *testing.T) {
	testHeader(t)
	shim, err := GetShim()
	HandleError(t, err)
	shim.StartServer()
	result, statusCode, err := HTTPGet(shim.Host())
	HandleError(t, err)
	shim.StopServer()
	assert.Equal(t, "Hello World!", result, `GET on Index should return "Hello World!"`)
	assert.Equal(t, 200, statusCode, `GET on Index should return 200 OK.`)
	testFooter(t)
}

// TestSunshinePath ensures that the test case given by Yoti works.
func TestSunshinePath(t *testing.T) {
	testHeader(t)
	shim, err := GetShim()
	HandleError(t, err)
	shim.StartServer()
	testCase := `
	{
		"roomSize" : [5, 5],
		"coords" : [1, 2],
		"patches" : [
			[1, 0],
			[2, 2],
			[2, 3]
		],
		"instructions" : "NNESEESWNWW"
	}
	`
	result, statusCode, err := HTTPPost(shim.Host(), testCase)
	HandleError(t, err)
	shim.StopServer()
	assert.Equal(t, 200, statusCode, `POST on Index should return a calculated response.`)
	assert.Equal(t, `{"coords":[1,3],"patches":1}`, result, `YOTI Example Response`)
	testFooter(t)
}

// TestRainyPath ensures that the if bad input is given (i.e. driving into a wall) the API responds properly.
func TestRainyPath(t *testing.T) {
	testHeader(t)
	shim, err := GetShim()
	HandleError(t, err)
	shim.StartServer()
	testCase := `
	{
		"roomSize" : [5, 5],
		"coords" : [1, 2],
		"patches" : [
			[1, 0],
			[2, 2],
			[2, 3]
			],
			"instructions" : "NNESEESWNWWWWWWWWWWWWWWWWWWWWWWWWWWWWWWW"
		}
		`
	result, statusCode, err := HTTPPost(shim.Host(), testCase)
	HandleError(t, err)
	shim.StopServer()
	assert.Equal(t, 200, statusCode, `POST on Index should return a calculated response.`)
	assert.Equal(t, `{"coords":[0,3],"patches":1}`, result, `Result should be equal`)
	testFooter(t)
}

// TestLargeRoomBehaviour ensures that if a really large room is passed in, the API responds on time.
func TestLargeRoomBehaviour(t *testing.T) {
	testHeader(t)
	shim, err := GetShim()
	HandleError(t, err)
	shim.StartServer()
	testCase := `
	{
		"roomSize" : [100, 100],
		"coords" : [50, 50],
		"patches" : [
			[0,0],
			[1,1],
			[2,2],
			[3,3],
			[4,4],
			[5,5],
			[6,6],
			[7,7],
			[8,8],
			[9,9],
			[10,10],
			[11,11],
			[12,12],
			[13,13],
			[14,14],
			[15,15],
			[16,16],
			[17,17],
			[18,18],
			[19,19],
			[20,20],
			[21,21],
			[22,22],
			[23,23],
			[24,24],
			[25,25],
			[26,26],
			[27,27],
			[28,28],
			[29,29],
			[30,30],
			[31,31],
			[32,32],
			[33,33],
			[34,34],
			[35,35],
			[36,36],
			[37,37],
			[38,38],
			[39,39],
			[40,40],
			[41,41],
			[42,42],
			[43,43],
			[44,44],
			[45,45],
			[46,46],
			[47,47],
			[48,48],
			[49,49],
			[50,50],
			[51,51],
			[52,52],
			[53,53],
			[54,54],
			[55,55],
			[56,56],
			[57,57],
			[58,58],
			[59,59],
			[60,60],
			[61,61],
			[62,62],
			[63,63],
			[64,64],
			[65,65],
			[66,66],
			[67,67],
			[68,68],
			[69,69],
			[70,70],
			[71,71],
			[72,72],
			[73,73],
			[74,74],
			[75,75],
			[76,76],
			[77,77],
			[78,78],
			[79,79],
			[80,80],
			[81,81],
			[82,82],
			[83,83],
			[84,84],
			[85,85],
			[86,86],
			[87,87],
			[88,88],
			[89,89],
			[90,90],
			[91,91],
			[92,92],
			[93,93],
			[94,94],
			[95,95],
			[96,96],
			[97,97],
			[98,98],
			[99,99],
			[100,100]
		],
		"instructions" : "NNNNNNNNNNNNNNNNNNNNNNNNNNNEEEEEEEEEEEEEEEEEEEEEEEEEEEEESSSSSSSSSSSSSSSSSSSSSSSSSSSSSSWWWWWWWWWWWWWWWWWWWWWWWWWWWNNNNNNNNNNNNNNNNNNNNNNNNNNEEEEEEEEEEEEEEEEEEEEEESSSSSSSSSSSSSSSSSSSSSSSSSSWWWWWWWWWWWWWWWWWWWWWWW"
	}
	`
	result, statusCode, err := HTTPPost(shim.Host(), testCase)
	HandleError(t, err)
	shim.StopServer()
	assert.Equal(t, 200, statusCode, `POST on Index should return a calculated response.`)
	assert.Equal(t, `{"coords":[51,47],"patches":4}`, result, `Result should be equal`)
	testFooter(t)
}
