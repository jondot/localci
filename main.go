package main

import(
  "flag"
	"github.com/go-fsnotify/fsnotify"
  "github.com/daviddengcn/go-colortext"
  "log"
  "time"
  "fmt"
  "os/exec"
  // or kingpin for nicer CLI
  // "gopkg.in/alecthomas/kingpin.v1"
)

func debounceChannel(interval time.Duration, input chan fsnotify.Event) chan fsnotify.Event {
  output := make(chan fsnotify.Event)
 
  go func() {
    var buffer fsnotify.Event
    var ok bool
    buffer, ok = <-input 
    if !ok {
      return
    }
    
    for {
      select {
      case buffer, ok = <-input:
        if !ok {
          return
        }
 
      case <-time.After(interval):
        output <- buffer
        buffer, ok = <-input
        if !ok {
          return
        }
      }
    }
  }()
  return output
}

var delay = flag.Int("n", 300, "Specify delay to ignore noisy events within an N millisecond window.")
var debug = flag.Bool("d", false, "Debug mode does not execute commands.")

func run(cmd string){
  defer ct.ResetColor()
  ct.ChangeColor(ct.Green, true, ct.None, false)
  log.Printf("Running: [%s]", cmd)

  out, err := exec.Command(cmd).Output()
  ct.ResetColor()
  fmt.Printf("%s\n", out)

  if err != nil {
      ct.ChangeColor(ct.Red, true, ct.None, false)
      fmt.Printf("Error: %s\n", err)
  }
  log.Printf("Ended: [%v]", cmd)
}

func main(){
  flag.Parse()
  if *delay < 50 {
    log.Fatalf("Delay must be greater than 50.");
  }

  var args = flag.Args()
  if len(args) == 0 {
    log.Fatalf("No args.")
  }
  log.Printf("Now watching: %v.", args)
  if(*debug){
    log.Printf("Debug mode enabled. No commands will really run.")
  }

  watcher, err := fsnotify.NewWatcher()
	if err != nil {
		log.Fatal(err)
	}
	defer watcher.Close()

	done := make(chan bool)


  debouncedEvents := debounceChannel(time.Duration(*delay)*time.Millisecond, watcher.Events)
	go func() {
		for {
			event := <-debouncedEvents
      if(*debug){
        log.Println("DEBUG:", event)

      } else {
        go func(){
          log.Println("RUN:", event)
          watcher.Remove(event.Name)
          run("./"+event.Name)
          watcher.Add(event.Name)
        }()
      }
		}
    log.Fatalf("ERROR: Cannot poll fs events anymore.")
	}()

	go func() {
		for {
			err := <-watcher.Errors
			log.Println("error:", err)
		}
    log.Fatalf("ERROR: Cannot poll errors events anymore.")
	}()

  for _, arg := range args {
    err = watcher.Add(arg)
    if err != nil {
      log.Fatal(err)
    }
  }

	<-done

}

