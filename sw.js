if(!self.define){let s,e={};const i=(i,n)=>(i=new URL(i+".js",n).href,e[i]||new Promise((e=>{if("document"in self){const s=document.createElement("script");s.src=i,s.onload=e,document.head.appendChild(s)}else s=i,importScripts(i),e()})).then((()=>{let s=e[i];if(!s)throw new Error(`Module ${i} didn’t register its module`);return s})));self.define=(n,r)=>{const l=s||("document"in self?document.currentScript.src:"")||location.href;if(e[l])return;let t={};const a=s=>i(s,l),o={module:{uri:l},exports:t,require:a};e[l]=Promise.all(n.map((s=>o[s]||a(s)))).then((s=>(r(...s),t)))}}define(["./workbox-2dfdff9b"],(function(s){"use strict";self.skipWaiting(),s.clientsClaim(),s.precacheAndRoute([{url:"assets/activity-service-66b44af8.js",revision:null},{url:"assets/activityService-64df3691.js",revision:null},{url:"assets/crop-black-072065eb.svg",revision:null},{url:"assets/destination-point-63b16bda.svg",revision:null},{url:"assets/ElevationGraph-0384c1d7.js",revision:null},{url:"assets/ElevationGraphView-68954a45.js",revision:null},{url:"assets/eye-off-black-537ff3d4.svg",revision:null},{url:"assets/fa-brands-400-003f1154.ttf",revision:null},{url:"assets/fa-brands-400-faae6fc0.woff2",revision:null},{url:"assets/fa-solid-900-886c8611.woff2",revision:null},{url:"assets/fa-solid-900-cea79b34.ttf",revision:null},{url:"assets/index-38c27766.css",revision:null},{url:"assets/main-d297a689.js",revision:null},{url:"assets/starting-point-c9a465c9.svg",revision:null},{url:"assets/wasmServices-8ea11c9f.js",revision:null},{url:"favicon.ico",revision:"1ba2ae710d927f13d483fd5d1e548c9b"},{url:"index.html",revision:"a69748de6e5a54a449d51474b350339f"},{url:"registerSW.js",revision:"1872c500de691dce40960bb85481de07"},{url:"wasm/activity-service.wasm",revision:"eaafdd625f245de3b953bf7c409cf5e0"},{url:"manifest.webmanifest",revision:"5349dda7e36ebe7251cdf3f3305d3f15"}],{}),s.cleanupOutdatedCaches(),s.registerRoute(new s.NavigationRoute(s.createHandlerBoundToURL("index.html"))),s.registerRoute(/^https:\/\/.*\.openstreetmap.org\/.*\.png$/,new s.CacheFirst,"GET")}));
