package godaemon

import (
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"strconv"
	"strings"
	"syscall"
)

/*
make the current daemon process PIDFile aware
in case of servers or so where another instance for same config shall result into error
empty param results to "/tmp/<binary-name>.pid"
*/
func PersistPID(pidFilename string) bool {
	if pidFilename == "" {
		pidFilename = PIDFilename()
	}

	_, current_status := ReadPIDFile(pidFilename)

	if !current_status {
		return UpdatePIDFile(pidFilename)
	}
	return false
}

/*
stop the daemon by killing the PIDs
empty param results to "/tmp/<binary-name>.pid"
*/
func KillPID(pidFilename string) bool {
	if pidFilename == "" {
		pidFilename = PIDFilename()
	}
	defer os.Remove(pidFilename)

	pid, current_status := ReadPIDFile(pidFilename)
	if current_status {
		if err := syscall.Kill(pid, 15); err == nil {
			return true
		}
	}
	return false
}

/*
status for stored PID or not
empty param results to "/tmp/<binary-name>.pid"
*/
func StatusPID(pidFilename string) {
	if pidFilename == "" {
		pidFilename = PIDFilename()
	}

	pid, current_status := ReadPIDFile(pidFilename)
	if current_status {
		fmt.Printf("Running. PID: '%d'", pid)
	} else if pid == -1 {
		fmt.Println("Stopped.")
	} else {
		fmt.Printf("Stopped. But PIDFile '%s' is present with PID '%d'.", pidFilename, pid)
	}
}

/*
return standard path to PIDFile for *nix nodes
*/
func PIDFilename() string {
	binary_path_token := strings.Split(os.Args[0], "/")
	binary_name := binary_path_token[len(binary_path_token)-1]
	return fmt.Sprintf("/tmp/%s.pid", binary_name)
}

/*
updates PID in given pidfile
if a daemon allows multiple instances, let it handle PIDFile naming
*/
func UpdatePIDFile(pidFilename string) bool {
	if pidFilename == "" {
		pidFilename = PIDFilename()
	}
	if _, err := os.Stat(pidFilename); err == nil {
		err := os.Remove(pidFilename)
		if err != nil {
			return false
		}
	}

	pid := strconv.Itoa(os.Getpid()) + "\n"
	return UpdateFile(pidFilename, pid)
}

/*
returns PID from pidfile and status of the process if stil running
*/
func ReadPIDFile(pidFilename string) (int, bool) {
	if pidFilename == "" {
		pidFilename = PIDFilename()
	}

	pid_bytes, _ := ioutil.ReadFile(pidFilename)
	pid_string := strings.TrimSpace(string(pid_bytes))
	pid, _ := strconv.Atoi(pid_string)
	if pid == 0 {
		return -1, false
	}

	if err := syscall.Kill(pid, 0); err == nil {
		return pid, true
	}
	return pid, false
}

/*
update given file content with given data
*/
func UpdateFile(filename string, filedata string) bool {
	f, err := os.Create(filename)
	if err != nil {
		fmt.Println(err)
		return false
	}
	n, err := io.WriteString(f, filedata)
	if err != nil {
		fmt.Println(n, err)
		return false
	}
	f.Close()
	return true
}
