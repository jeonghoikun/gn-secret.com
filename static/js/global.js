"use strict";
/// <reference path="./@types.d.ts" />
// dev
let log = console.log;
// selectors
function $(q) { return document.querySelector(q); }
function $$(q) { return document.querySelectorAll(q); }
class Luxon {
    static now() { return luxon.DateTime.now(); }
    static fromISO(iso) { return luxon.DateTime.fromISO(iso); }
    static seoul(iso) {
        let zone = "Asia/Seoul";
        if (iso == undefined) {
            return this.now().setZone(zone);
        }
        return this.fromISO(iso).setZone(zone);
    }
}
class Time {
    constructor(iso) {
        let seoul = Luxon.seoul(iso);
        let c = seoul.c;
        this.year = Time.itoa(c.year);
        this.month = Time.itoa(c.month);
        this.day = Time.itoa(c.day);
        this.weekday = Time.numToWeekdayKO(seoul.weekday);
        this.hour = Time.itoa(c.hour);
        this.minute = Time.itoa(c.minute);
        this.second = Time.itoa(c.second);
    }
    static itoa(n) {
        let nString = String(n);
        if (n < 10) {
            return "0" + nString;
        }
        return nString;
    }
    static numToWeekdayKO(n) {
        switch (n) {
            case 1:
                return "월";
            case 2:
                return "화";
            case 3:
                return "수";
            case 4:
                return "목";
            case 5:
                return "금";
            case 6:
                return "토";
            case 7:
                return "일";
            default:
                return "<err>";
        }
    }
    format(s) {
        switch (s) {
            case "base":
                return `${this.year}-${this.month}-${this.day} ${this.hour}:${this.minute}`;
            default:
                return "ERROR: Invalid format";
        }
    }
    static sleep(ms) {
        return new Promise((rsv) => { setTimeout(rsv, ms); });
    }
}
// Constants
Time.second = 1000;
