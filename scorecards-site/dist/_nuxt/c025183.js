(window.webpackJsonp=window.webpackJsonp||[]).push([[7],{367:function(e,n,t){"use strict";t.r(n);var r=t(109);t(56),t(12),t(24),t(222),t(42),t(26),t(45),t(33),t(32),t(25),t(46),t(47),t(34);function o(e,n){var t="undefined"!=typeof Symbol&&e[Symbol.iterator]||e["@@iterator"];if(!t){if(Array.isArray(e)||(t=function(e,n){if(!e)return;if("string"==typeof e)return c(e,n);var t=Object.prototype.toString.call(e).slice(8,-1);"Object"===t&&e.constructor&&(t=e.constructor.name);if("Map"===t||"Set"===t)return Array.from(e);if("Arguments"===t||/^(?:Ui|I)nt(?:8|16|32)(?:Clamped)?Array$/.test(t))return c(e,n)}(e))||n&&e&&"number"==typeof e.length){t&&(e=t);var i=0,r=function(){};return{s:r,n:function(){return i>=e.length?{done:!0}:{done:!1,value:e[i++]}},e:function(e){throw e},f:r}}throw new TypeError("Invalid attempt to iterate non-iterable instance.\nIn order to be iterable, non-array objects must have a [Symbol.iterator]() method.")}var o,f=!0,l=!1;return{s:function(){t=t.call(e)},n:function(){var e=t.next();return f=e.done,e},e:function(e){l=!0,o=e},f:function(){try{f||null==t.return||t.return()}finally{if(l)throw o}}}}function c(e,n){(null==n||n>e.length)&&(n=e.length);for(var i=0,t=new Array(n);i<n;i++)t[i]=e[i];return t}var f={props:{selector:{type:String,default:""},once:{type:Boolean,default:!1}},data:function(){return{items:[],onScreen:!1}},watch:{onScreen:function(e){var n=this;this.once&&e&&this.$nuxt.$on("observer.observed",(function(e){var t,r=o(n.items);try{for(r.s();!(t=r.n()).done;){var c=t.value.el;e.unobserve(c)}}catch(e){r.e(e)}finally{r.f()}}))}},mounted:function(){var e=this;this.items=Object(r.a)(this.selector?document.querySelectorAll(this.selector):[this.$el]).map((function(e){return{el:e,onScreen:!1}})),this.items.length&&(this.$nuxt.$on("observer.observed",(function(n){n.forEach((function(n){var t=e.items.find((function(e){var t=e.el;return n.target===t}));t&&(t.onScreen=n.isIntersecting)})),e.updateOnScreen()})),this.$nuxt.$on("observer.created",(function(n){var t,r=o(e.items);try{for(r.s();!(t=r.n()).done;){var c=t.value.el;n.observe(c)}}catch(e){r.e(e)}finally{r.f()}})),this.updateOnScreen())},methods:{updateOnScreen:function(){this.onScreen=!!this.items.find((function(e){return e.onScreen}))}},render:function(){return this.$scopedSlots.default({onScreen:this.onScreen})}},l=t(20),component=Object(l.a)(f,undefined,undefined,!1,null,null,null);n.default=component.exports}}]);