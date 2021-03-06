///////////////////////////////////////////////////////////////////////////////////////////////////////////////////////

package main

import (
	"fmt"
	"os"
	"bufio"
	"log"
	"io"
)

type voter struct{
	name string
	id int8
	vote string
}
type party struct{
	PTI int64
	PMLN int64
	PPPP int64
	MQM int64
}

func Scanfile() {                        // This function Scans the last saved files
	file2, _ := os.Open("result.txt")    //Opening last saved file
	scanner := bufio.NewScanner(file2)
	for scanner.Scan() {             
        fmt.Println(scanner.Text())       //Printing file
    }
}

func Error () *os.File {							// This Function creates a new file if it doesnt exist
	nill:=0
         
                 file, err := os.Create("data.txt")
					if err != nil {
						log.Fatal("Cannot create file", err)
					}
					defer file.Close()
					j:=0
					for j<5{
						fmt.Fprintf(file,"%d\r\n",nill)
						fmt.Println("Please Restart the software!")
						j++
					}
        
		 return file
}


func CreatedFile () *os.File{                 // This Functions creates new data file
	
	dfile, err := os.Create("data.txt")
    if err != nil {
        log.Fatal("Cannot create file", err)
    }
	return dfile
}

func CreatrFile () *os.File{                 // This Function creates new result File
	rfile, err := os.Create("result.txt")
    if err != nil {
        log.Fatal("Cannot create file", err)
    }
	return rfile
	}

func UserInput(m int64, standing party, rfile *os.File, dfile *os.File){                   // This Function prompts the user to input data
	caster := voter {"",0,""}   //intializing structure
	for m<1000{
		fmt.Println("\n Welcome to the Electronic Voting System.\n")
	
		fmt.Print("\n Please Enter your full name.: \n")     // Voter data entry
		reader := bufio.NewReader(os.Stdin)
		text, _ := reader.ReadString('\n')
		caster.name=text
		fmt.Println("\nEnter Your ID")
		fmt.Scanln(&caster.id)
		fmt.Println("\n Please Select the party you want to vote")
		fmt.Println("1) PTI",
					"\n2) PMLN",
					"\n3) PPPP",
					"\n4) MQM")
		fmt.Scanln(&caster.vote)
	
		fmt.Println("\nPlease Confirm Your Data: ")         // Data Confirmation
		fmt.Println("\nYour Name:", caster.name,
					"\nYour ID:", caster.id,
					"\nYour selected party:", caster.vote)
		
		switch{												//Total vote calculator
			case caster.vote=="PTI":
				standing.PTI++
				m++
			case caster.vote=="PMLN":
				standing.PMLN++
				m++
			case caster.vote=="PPPP":
				standing.PPPP++
				m++
			case caster.vote=="MQM":
				standing.MQM++
				m++
		}
		
		
		var des int
		fmt.Println("\nEnter 1 to continue...")
		fmt.Scanln(&des)   //Enter 2 to exit and see results
		if des==2{ break }
	}
	
	FileWriter(rfile,dfile,m,standing)            // Function call to write data in result and data file
	Result(standing)                              // Function to Display results
}

func FileWriter (rfile,dfile *os.File, m int64, standing party) {                                   // This function writes the results in the file
	fmt.Fprintf(rfile, "Total Voters %d\r\n",m)              // Log file writer
		fmt.Fprintf(dfile,"%d\r\n",m)
		
		fmt.Fprintf(rfile,"PTI: %d\r\n",standing.PTI)
		fmt.Fprintf(dfile,"%d\r\n",standing.PTI)
		
		fmt.Fprintf(rfile,"PMLN: %d\r\n",standing.PMLN)
		fmt.Fprintf(dfile,"%d\r\n",standing.PMLN)
		
		fmt.Fprintf(rfile,"PPPP: %d\r\n",standing.PPPP)
		fmt.Fprintf(dfile,"%d\r\n",standing.PPPP)
		
		fmt.Fprintf(rfile,"MQM: %d\r\n",standing.MQM)
		fmt.Fprintf(dfile,"%d\r\n",standing.MQM)
}

func Result (standing party) {
	fmt.Println("\nVoting Result:\n")                             // Total votes
	fmt.Println("Votes for PTI: ", standing.PTI,
				"\nVotes for PMLN: ",standing.PMLN,
				"\nVotes for PPPP: ", standing.PPPP,
				"\nVotes for MQM: ", standing.MQM, "\n\n")
}

func main(){           // Main Functions

	var m int64
	
	Scanfile()                       // Function call to Scan the files
	file, err := os.Open("data.txt")
	if err != nil {
			Error()
		}
    var perline int64
	var nums []int64
	
	for {

                 _, err := fmt.Fscanf(file, "%d\n", &perline) // give a patter to scan

                 if err != nil {

                         if err == io.EOF {
                                 break // stop reading the file
                         }
                 }

                 nums = append(nums, perline)
				 
         }
	
	standing := party{nums[1],nums[2],nums[3],nums[4]}
	m = nums[0]
	
	dfile:=CreatedFile()        // Function call to create new data file
	rfile:=CreatrFile()         // Function call to create new result file
	UserInput(m,standing,rfile,dfile)    //Function call to prompt the user to input data
	
}
