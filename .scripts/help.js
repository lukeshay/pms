#!/usr/bin/env node

import fs from "node:fs";

const BOLD_ON = "\x1b[1m";
const BOLD_OFF = "\x1b[22m";

/** @type {Record<"scripts", Record<string, string>>} */
const packageJson = JSON.parse(fs.readFileSync("./package.json"));

Object.entries(packageJson.scripts)
  .filter(([key]) => key.includes(" - comment"))
  .forEach(([key, value]) => {
    const scriptName = key.replace(" - comment", "");
    console.log(`${BOLD_ON}${scriptName}:${BOLD_OFF} ${value}`);
  });
