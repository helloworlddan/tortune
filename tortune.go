package main

import (
	"flag"
	"fmt"
	"math/rand"
	"time"
)

const version = "0.0.3"

var (
	db = [...]string{
		"Zero-Trust means everybody gets access to Ring-0",
		"Localhost?! More like loserhost!",
		"alias yolo=\"sudo\"",
		"Friends don't let friends compile locally.",
		"My other supercomputer is Google Cloud.",
		"az·ure (/ˈaZHər/): bright blue in color like a cloudless sky.",
		"What do clouds wear under their pants?\nThunderwear.",
		"With the increase in cloud computing, now is not a good time to invest in solar energy.",
		"Cloud Computing? Computers can fly now?",
		"There's a band called 1023MB. Ever heard of them? They haven't had a gig yet.",
		"If Bill Gates had a penny for every time I had to reboot my computer... oh wait, he does.",
		"CAPS LOCK: Preventing Logins Since 1980.",
		"I needed a password eight characters long so I picked SnowWhiteandtheSevenDwarves.",
		"What do you call it when computer science majors make fun of each other? Cyber boolean!",
		"I love the F5 key. It's just so refreshing.",
		"Normal people use their children's names to set their email passwords. Elon Musk uses his email password (X Æ A-12) to name his baby.",
		"Do not use 'beef stew' as a computer password. It is not stroganoff.",
		"99 little bugs in the code,\n99 little bugs,\nTake one down, patch it around,\n117 little bugs in the code.",
		"My cousin just got fired from the keyboard factory. They said he just wasn't putting in enough shifts!",
		"So, the other day I started to whisper and my wife asked why I was whispering? I told her I didn't want Mark Zuckerberg to hear us.",
		"I laughed.\nMy wife laughed.\nAlexa laughed.\nSiri laughed.",
		"Where did the cybersecurity team go the last few days?\nThey ran-som-ware.",
		"Why doesn't Superman fight cybercrime?\nHe's afraid of Krypto Currency!",
		"What did the hacker's out of office message say?\nGone phishing.",
		"What do you call a turtle that surfs the Dark Web?\nA TORtoise.",
		"What do you call a group of math and science geeks at a party?\nSocial engineers!",
		"Cybersecurity is like an Onion.\nThere's layers, and at some point you start to cry.",
		"The 'S' in IoT stands for Security.",
		"Even though when I show people my password they say it's secure, I still get hacked sometimes.",
		"I used to tell a joke about UDP at parties, but I quit.\nI could never tell if anyone got it.",
	}
)

func main() {
	versionToggle := flag.Bool("version", false, "show version info")
	flag.Parse()

	if *versionToggle == true {
		fmt.Printf("tortune version %s\n", version)
		return
	}

	rand.Seed(time.Now().Unix())
	fmt.Println(db[rand.Intn(len(db))])
}
