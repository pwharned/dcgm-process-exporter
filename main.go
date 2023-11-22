package main

import (
	"fmt"
	"log"
	"time"
	"github.com/NVIDIA/go-nvml/pkg/nvml"
	"net/http"
	"strings"
)

func main() {
 
    // API routes
    http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
        fmt.Fprintf(w, "Hello world from GfG")
    })
    http.HandleFunc("/metrics", func(w http.ResponseWriter, r *http.Request) {
        fmt.Fprintf(w, strings.Join(getMetrics(), "\n"))
    })
 
    port := ":5000"
    fmt.Println("Server is running on port" + port)
 
    // Start server on port specified above
    log.Fatal(http.ListenAndServe(port, nil))
}


	//Pid       uint32
	//TimeStamp uint64
	//SmUtil    uint32
	//MemUtil   uint32
	//EncUtil   uint32
	//DecUtil   uint32




func getMetrics() []string {

	
	

	ret := nvml.Init()
	if ret != nvml.SUCCESS {
		log.Fatalf("Unable to initialize NVML: %v", nvml.ErrorString(ret))
	}
	defer func() {
		ret := nvml.Shutdown()
		if ret != nvml.SUCCESS {
			log.Fatalf("Unable to shutdown NVML: %v", nvml.ErrorString(ret))
		}
	}()

	count, ret := nvml.DeviceGetCount()

	var stringResult  []string//[]nvml.ProcessUtilizationSample


	if ret != nvml.SUCCESS {
		log.Fatalf("Unable to get device count: %v", nvml.ErrorString(ret))
	}

	for i := 0; i < count; i++ {
		device, ret := nvml.DeviceGetHandleByIndex(i)
		if ret != nvml.SUCCESS {
			log.Fatalf("Unable to get device at index %d: %v", i, nvml.ErrorString(ret))
		}

		uuid, ret := device.GetUUID()
		if ret != nvml.SUCCESS {
			log.Fatalf("Unable to get uuid of device at index %d: %v", i, nvml.ErrorString(ret))
		}
		
		fmt.Printf("Device id %v",uuid)

		timestamp := uint64(time.Now().Unix())

		processutilizationsample, ret := nvml.DeviceGetProcessUtilization(device, timestamp) 
		var stringArray []string

		if ret != nvml.SUCCESS {
		log.Fatalf("Unable to get Device Process Utilization %v", nvml.ErrorString(ret) )
		}
		
		for j:=0; j<len(processutilizationsample); j++{

			pid :=processutilizationsample[j].Pid
			
			//fmt.Printf(fmt.Sprintf(" Hello i am an int %v\n",pid))
		stringArray = append(stringArray[:],fmt.Sprintf(`PROCESS_UTILIZATION{pid="%v"} %v`,pid, processutilizationsample[j].SmUtil))

		
	}
	
	stringResult = append(stringResult[:], stringArray[:]...)


}
return stringResult
}
