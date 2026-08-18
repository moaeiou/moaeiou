#!/usr/bin/env node
"use strict";

import { readFileSync, writeFileSync } from "fs";

const [src, dest] = process.argv.slice(2);

const minified = readFileSync(src, "utf8")
  .replace(/<!--[\s\S]*?-->/g, "")
  .replace(/<[^>]+>/g, (tag) =>
    tag.replace(/\s+/g, " ").replace(/< /, "<").replace(/\s+>/g, ">"),
  )
  .replace(/>\s+</g, "><")
  .replace(/>([^<]+)</g, (_, text) => `>${text.replace(/\s+/g, " ")}<`)
  .trim();

writeFileSync(dest, minified);