package uptime

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"time"
)

const (
	UPTIME_PATH  = "/proc/uptime"
	LOADAVG_PATH = "/proc/loadavg"
)

// Uptime represents performance metric computed by system.
type Uptime struct {
	Now                  time.Time
	BootTime             float64
	LAvg1, LAvg5, LAvg15 float64
}

// New is a constructor for Uptime
func New() (*Uptime, error) {
	boot, err := readUptime()
	if err != nil {
		return nil, err
	}

	lavg, err := readLAvg()
	if err != nil {
		return nil, err
	}

	return &Uptime{
		Now:      time.Now(),
		BootTime: boot,
		LAvg1:    lavg[0],
		LAvg5:    lavg[1],
		LAvg15:   lavg[2],
	}, nil
}

func (uptm *Uptime) calcBootExactTime() (int, int) {
	d := time.Duration(uptm.BootTime) * time.Second
	min := int(d.Minutes())
	return min / 60, min % 60
}

func (uptm *Uptime) calcBootSince() time.Time {
	d := time.Duration(uptm.BootTime) * time.Second
	return uptm.Now.Add(-d)

}

// Print prints the entire performance metrics in one line.
func (uptm *Uptime) Print() {
	hours, mins := uptm.calcBootExactTime()

	w := bufio.NewWriter(os.Stdout)
	w.WriteString(" ")
	w.WriteString(uptm.Now.Format("15:04:05"))
	w.WriteString(" ")
	w.WriteString(fmt.Sprintf("up  %d:%2d,", hours, mins))
	w.WriteString("  ")
	w.WriteString("1 user,")
	w.WriteString("  ")
	w.WriteString(fmt.Sprintf("load average: %2.2f, %2.2f, %2.2f", uptm.LAvg1, uptm.LAvg5, uptm.LAvg15))
	w.WriteString("\n")

	w.Flush()
}

// PrettyPrint prints the current uptime with hours and minutes.
func (uptm *Uptime) PrettyPrint() {
	hours, mins := uptm.calcBootExactTime()

	w := bufio.NewWriter(os.Stdout)
	w.WriteString(fmt.Sprintf("up %d hours, %d minutes", hours, mins))
	w.WriteString("\n")

	w.Flush()
}

// SincePrint prints the time since booting in pretty format.
func (uptm *Uptime) SincePrint() {
	since := uptm.calcBootSince()

	w := bufio.NewWriter(os.Stdout)
	w.WriteString(since.Format("2006-01-02 15:04:05"))
	w.WriteString("\n")

	w.Flush()
}

func readUptime() (boot float64, err error) {
	f, err := os.Open(UPTIME_PATH)
	if err != nil {
		err = fmt.Errorf("failed to open file: %s\n  %s", UPTIME_PATH, err)
		return
	}
	defer f.Close()

	sc := bufio.NewScanner(f)
	sc.Split(bufio.ScanWords)

	if sc.Scan() {
		boot, err = strconv.ParseFloat(sc.Text(), 64)
		if err != nil {
			return
		}
	}

	return
}

func readLAvg() (lavg [3]float64, err error) {
	f, err := os.Open(LOADAVG_PATH)
	if err != nil {
		err = fmt.Errorf("failed to open file: %s\n  %s", LOADAVG_PATH, err)
		return
	}
	defer f.Close()

	sc := bufio.NewScanner(f)
	sc.Split(bufio.ScanWords)

	for i := 0; i < 3 && sc.Scan(); i++ {
		lavg[i], err = strconv.ParseFloat(sc.Text(), 64)
		if err != nil {
			err = fmt.Errorf("lavg%d read failed: %s", i, err)
			return
		}
	}

	return
}
