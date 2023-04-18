package tortune

import (
	"math/rand"
	"time"
)

const Version = "0.0.13"

var (
	db = [...]string{
		"Chrome is here the heart is",
		"Chrome Sweet Chrome",
		"There is no place like Chrome",
		"['🥚','🐔'].shuffle.pop",
		"A: Hey, do you know my friend Jason?\nB: You mean Jason Marshall?!\nYeah, he's super serial.",
		"Microsoft has really gotten good at making hardware. I colloquially refer to this phenomenon as Microhard.",
		"Interviewer: What is your biggest strength?\nMe: I am an expert in machine learning.\nInterviewer: What’s 6 + 10?\nMe: Zero.\nInterviewer: Nowhere near. It’s 16.\nMe: Ok, It’s 16.\nInterviewer: What is 10 + 20?\nMe: It’s 16.",
		"I was surprised when my niece said she learned R at school yesterday, and then I remembered she's 4 and she meant the letter.",
		"Q: Why did the naive Bayesian suddenly feel patriotic when he heard fireworks?\nA: It assumed independence.",
		"Movie exec: 'Pitch me!'\nMe: 'It’s a movie about high school girls trying to figure out what clique they belong in. They move from clique to clique and eventually stop when they minimize their differences. It’s called K-Means girls.'",
		"Your ML models are looking real fit, they must do a lot of weighted training.",
		"Chuck Norris can take a screenshot of his blue screen.",
		"Q: Why can’t data engineers become hat makers?\nA: They can only guarantee two thirds of a CAP!",
		"Nothing good can come from 2989 witches casting a hex; it's always 0xBAD.",
		"Good programmers copy; great programmers paste.",
		"Programmers store all their dad jokes in a dad-a-base.",
		"Unfortunately these jokes only work if you git them.",
		"Next year will be the year of Linux on the desktop!",
		"Q: Why did the functions stop calling each other?\nA: Because they had constant arguments.",
		"There are 2 hard problems in computer science: caching, naming, and off-by-1 errors.",
		"UNIX is user friendly. It’s just very particular about who its friends are.",
		"Two threads walk into a bar. The barkeeper looks up and yells, “Hey, I want don’t any conditions race like time last!”",
		"ASCII stupid question, get a stupid ANSI.",
		"Real programmers count from 0.",
		"Debugging: Being the detective in a movie where you are also the murderer.",
		"Don't assume; it makes an ASS out of U and ME.",
		"Zero-Trust means everybody gets access to Ring-0.",
		"Localhost?! More like loserhost!",
		"alias yolo=\"sudo !!\"",
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
		"Q: How many Prolog programmers does it take to change a lightbulb?\nA: Yes.",
		"What can you do if you cannot push your git changes?\nUse the '--force', Luke",
		"How do you get the code for the bank vault?\nYou checkout their branch.",
		"Why doesn't Hollywood make more Big Data movies?\nNoSQL",
		"How do programming pirates pass method parameters?\nVarrrrarrrgs.",
		"How did pirates collaborate before computers?\nPier to pier networking.",
		"['hip','hip'] // (hip hip array!)",
		"Halloween and Christmas are so confusing, because 31 OCT == 25 DEC.",
		"UDP bar is packet going into a",
		"Hi. I'd like to hear a TCP joke.\nHello, would you like to hear a TCP joke?\nYes, I would like to hear a TCP joke.\nOK, I will tell you a TCP joke.\nOk, I will hear a TCP joke.\nAre you ready to hear a TCP joke?\nYes. I am ready to hear a TCP joke.\nOk, I am about to send the TCP joke. It will last 10 seconds, it has two characters, it does not have a setting, it ends with a punchline.\nOk, I am ready to receive your TCP joke that will last 10 seconds, has two...\nI'm sorry, your connection has timed out.\nHi, I would like to hear a TCP joke?",
	}
)

func HitMe() string {
	rand.Seed(time.Now().Unix())

	return db[rand.Intn(len(db))]
}
