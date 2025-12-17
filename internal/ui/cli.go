package ui

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/G1P0/hopless/internal/domain"
	"github.com/G1P0/hopless/internal/engine"
)

type CLI struct {
	world  domain.World
	reader *bufio.Reader
}

func NewCLI() *CLI {
	return &CLI{
		world:  domain.World{},
		reader: bufio.NewReader(os.Stdin),
	}
}

func (c *CLI) Run() {
	fmt.Println("HOPLESS")
	fmt.Println("no route. no hope.")

	c.world.Links = []domain.Link{
		{From: domain.Client, To: domain.Router},
		{From: domain.Router, To: domain.Server},
	}

	for {
		fmt.Println("\n--- MENU ---")
		fmt.Println("1) show mission")
		fmt.Println("2) show rules")
		fmt.Println("3) add rule")
		fmt.Println("4) delete rule")
		fmt.Println("5) ping")
		fmt.Println("6) check mission")
		fmt.Println("7) reset world")
		fmt.Println("0) exit")

		switch c.readLine("> ") {
		case "1":
			fmt.Println(engine.MissionText())
		case "2":
			c.showRules()
		case "3":
			c.addRule()
		case "4":
			c.delRule()
		case "5":
			c.ping()
		case "6":
			ok, reasons := engine.MissionComplete(c.world)
			for _, s := range reasons {
				fmt.Println(s)
			}
			if ok {
				fmt.Println("\nMISSION COMPLETE ✅")
			} else {
				fmt.Println("\nMISSION FAILED ❌")
			}
		case "7":
			c.world = domain.World{}
			fmt.Println("World reset. Back to DENY by default.")
		case "0":
			fmt.Println("bye.")
			return
		default:
			fmt.Println("Unknown command.")
		}
	}
}

func (c *CLI) readLine(prompt string) string {
	fmt.Print(prompt)
	line, _ := c.reader.ReadString('\n')
	return strings.TrimSpace(line)
}

func parseNode(s string) (domain.Node, bool) {
	switch strings.ToLower(strings.TrimSpace(s)) {
	case "client":
		return domain.Client, true
	case "router":
		return domain.Router, true
	case "server":
		return domain.Server, true
	default:
		return "", false
	}
}

func (c *CLI) showRules() {
	if len(c.world.Rules) == 0 {
		fmt.Println("No rules. Everything is DENY by default.")
		return
	}
	fmt.Println("Rules (last match wins):")
	for i, r := range c.world.Rules {
		action := "DENY"
		if r.Allow {
			action = "ALLOW"
		}
		fmt.Printf("  #%d  %s  %s:%d -> %s\n", i, action, r.Src, r.Port, r.Dst)
	}
}

func (c *CLI) addRule() {
	fromStr := c.readLine("from (client/router/server): ")
	from, ok := parseNode(fromStr)
	if !ok {
		fmt.Println("Unknown node:", fromStr)
		return
	}
	toStr := c.readLine("to   (client/router/server): ")
	to, ok := parseNode(toStr)
	if !ok {
		fmt.Println("Unknown node:", toStr)
		return
	}

	portStr := c.readLine("port (number, 0=ANY): ")
	port, err := strconv.Atoi(strings.TrimSpace(portStr))
	if err != nil {
		fmt.Println("Not a number:", portStr)
		return
	}
	if port < 0 || port > 65535 {
		fmt.Println("Bad port:", port)
		return
	}

	allowStr := strings.ToLower(c.readLine("allow? (yes/no): "))
	var allow bool
	switch allowStr {
	case "y", "yes":
		allow = true
	case "n", "no":
		allow = false
	default:
		fmt.Println("Expected yes/no.")
		return
	}

	c.world.Rules = append(c.world.Rules, domain.Rule{Src: from, Dst: to, Allow: allow, Port: port})
	fmt.Println("Rule added.")
}

func (c *CLI) delRule() {
	if len(c.world.Rules) == 0 {
		fmt.Println("No rules to delete.")
		return
	}
	c.showRules()
	idxStr := c.readLine("delete rule # (number): ")
	var idx int
	_, err := fmt.Sscanf(idxStr, "%d", &idx)
	if err != nil || idx < 0 || idx >= len(c.world.Rules) {
		fmt.Println("Bad index.")
		return
	}
	c.world.Rules = append(c.world.Rules[:idx], c.world.Rules[idx+1:]...)
	fmt.Println("Rule deleted.")
}

func (c *CLI) ping() {
	fromStr := c.readLine("ping from (client/router/server): ")
	from, ok := parseNode(fromStr)
	if !ok {
		fmt.Println("Unknown node:", fromStr)
		return
	}

	toStr := c.readLine("ping to   (client/router/server): ")
	to, ok := parseNode(toStr)
	if !ok {
		fmt.Println("Unknown node:", toStr)
		return
	}

	portStr := c.readLine("port (number, 0=ANY): ")
	var port int
	if _, err := fmt.Sscanf(portStr, "%d", &port); err != nil || port < 0 || port > 65535 {
		fmt.Println("Bad port.")
		return
	}

	ok2, why := engine.CanReachRouted(c.world, engine.Query{
		From: from,
		To:   to,
		Port: port,
	})

	if ok2 {
		fmt.Printf("PING %s -> %s port=%d: OK (%s)\n", from, to, port, why)
	} else {
		fmt.Printf("PING %s -> %s port=%d: FAIL (%s)\n", from, to, port, why)
	}
}
