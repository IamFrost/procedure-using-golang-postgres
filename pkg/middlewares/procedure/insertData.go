package procedure

import(
	"github.com/IamFrost/test5/pkg/utils"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

// InsertDataProCall will call insert_data procedure
// that will insert 2 values in database table - tbl
func InsertDataProCall(c *gin.Context){
	c.Writer.Header().Set("Access-Control-Allow-Origin", "*")

	dbpool := utils.CreateConnection()
	defer dbpool.Close()

	integer1FromRouter := c.Param("integer1")
	integer2FromRouter := c.Param("integer2")
	
	myQuery := `CALL insert_data($1, $2)`

	result, err := dbpool.Exec(c, myQuery, integer1FromRouter, integer2FromRouter)

	if err != nil {
		panic(err)
	} else {
		logrus.Debugf("From InsertDataProCall :  %v", result)
	}

}