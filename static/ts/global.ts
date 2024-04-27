/// <reference path="./@types.d.ts" />

// dev
let log: Function = console.log

// selectors
function $(q: string): unknown  { return document.querySelector(q) }
function $$(q: string): unknown { return document.querySelectorAll(q) }

class Luxon {
	private static now(): any                { return luxon.DateTime.now() }
	private static fromISO(iso: string): any { return luxon.DateTime.fromISO(iso) }

	public  static seoul(iso?: string): any {
		let zone = "Asia/Seoul"
		if (iso == undefined) {
			return this.now().setZone(zone)
		}
		return this.fromISO(iso).setZone(zone)
	}
}

class Time {
	// Constants
	public static second = 1000
	// Vars
	public year:    string
	public month:   string
	public day:     string
	public weekday: string
	public hour:    string
	public minute:  string
	public second:  string
	public constructor(iso?: string) {
		let seoul: any = Luxon.seoul(iso)
		let c:     any = seoul.c
		this.year    = Time.itoa(c.year)
		this.month   = Time.itoa(c.month)
		this.day     = Time.itoa(c.day)
		this.weekday = Time.numToWeekdayKO(seoul.weekday)
		this.hour    = Time.itoa(c.hour)
		this.minute  = Time.itoa(c.minute)
		this.second  = Time.itoa(c.second)
	}
	private static itoa(n: number): string {
		let nString: string = String(n)
		if (n < 10) {
			return "0" + nString
		}
		return nString
	}
	private static numToWeekdayKO(n: number): string {
		switch (n) {
		case 1:
			return "월"
		case 2:
			return "화"
		case 3:
			return "수"
		case 4:
			return "목"
		case 5:
			return "금"
		case 6:
			return "토"
		case 7:
			return "일"
		default:
			return "<err>"
		}
	}
	public format(s: string): string {
		switch (s) {
		case "base":
			return `${this.year}-${this.month}-${this.day} ${this.hour}:${this.minute}`
		default:
			return "ERROR: Invalid format"
		}
	}
	public static sleep(ms: number): Promise<void> {
		return new Promise((rsv: any) => { setTimeout(rsv, ms) })
	}
}
