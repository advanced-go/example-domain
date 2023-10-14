package guidance

import "fmt"

func ExamplePutSLO() {
	PutSLO(SLO{Controller: "percentile", Threshold: "99/1s", StatusCodes: ""})
	fmt.Printf("test: PutSLO() -> %v\n", list)

	PutSLO(SLO{Controller: "status-codes", Threshold: "10%", StatusCodes: "500,504"})
	fmt.Printf("test: PutSLO() -> %v\n", list)

	PutSLO(SLO{Controller: "percentile", Threshold: "95/500ms", StatusCodes: ""})
	fmt.Printf("test: PutSLO() -> %v\n", list)

	//Output:
	//test: PutSLO() -> [{percentile 99/1s }]
	//test: PutSLO() -> [{percentile 99/1s } {status-codes 10% 500,504}]
	//test: PutSLO() -> [{percentile 95/500ms } {status-codes 10% 500,504}]

}

func ExampleGetSLOByController() {
	PutSLO(SLO{Controller: "percentile", Threshold: "99/1s", StatusCodes: ""})
	PutSLO(SLO{Controller: "status-codes", Threshold: "10%", StatusCodes: "500,504"})

	ctrl := ""
	s := GetSLOByController(ctrl)
	fmt.Printf("test: GetSLOByController(%s) -> %v\n", ctrl, s)

	ctrl = "percentile"
	s = GetSLOByController(ctrl)
	fmt.Printf("test: GetSLOByController(%s) -> %v\n", ctrl, s)

	ctrl = "status-codes"
	s = GetSLOByController(ctrl)
	fmt.Printf("test: GetSLOByController(%s) -> %v\n", ctrl, s)

	PutSLO(SLO{Controller: "percentile", Threshold: "95/500ms", StatusCodes: ""})

	ctrl = "percentile"
	s = GetSLOByController(ctrl)
	fmt.Printf("test: GetSLOByController(%s) -> %v\n", ctrl, s)

	//Output:
	//test: GetSLOByController() -> <nil>
	//test: GetSLOByController(percentile) -> &{percentile 99/1s }
	//test: GetSLOByController(status-codes) -> &{status-codes 10% 500,504}
	//test: GetSLOByController(percentile) -> &{percentile 95/500ms }

}
