(function(){/*

 Copyright The Closure Library Authors.
 SPDX-License-Identifier: Apache-2.0
*/
var r;function aa(a){var b=0;return function(){return b<a.length?{done:!1,value:a[b++]}:{done:!0}}}
var ba="function"==typeof Object.defineProperties?Object.defineProperty:function(a,b,c){if(a==Array.prototype||a==Object.prototype)return a;a[b]=c.value;return a};
function ca(a){a=["object"==typeof globalThis&&globalThis,a,"object"==typeof window&&window,"object"==typeof self&&self,"object"==typeof global&&global];for(var b=0;b<a.length;++b){var c=a[b];if(c&&c.Math==Math)return c}throw Error("Cannot find global object");}
var da=ca(this);function t(a,b){if(b)a:{for(var c=da,d=a.split("."),e=0;e<d.length-1;e++){var f=d[e];if(!(f in c))break a;c=c[f]}d=d[d.length-1];e=c[d];f=b(e);f!=e&&null!=f&&ba(c,d,{configurable:!0,writable:!0,value:f})}}
t("Symbol",function(a){function b(e){if(this instanceof b)throw new TypeError("Symbol is not a constructor");return new c("jscomp_symbol_"+(e||"")+"_"+d++,e)}
function c(e,f){this.h=e;ba(this,"description",{configurable:!0,writable:!0,value:f})}
if(a)return a;c.prototype.toString=function(){return this.h};
var d=0;return b});
t("Symbol.iterator",function(a){if(a)return a;a=Symbol("Symbol.iterator");for(var b="Array Int8Array Uint8Array Uint8ClampedArray Int16Array Uint16Array Int32Array Uint32Array Float32Array Float64Array".split(" "),c=0;c<b.length;c++){var d=da[b[c]];"function"===typeof d&&"function"!=typeof d.prototype[a]&&ba(d.prototype,a,{configurable:!0,writable:!0,value:function(){return ea(aa(this))}})}return a});
function ea(a){a={next:a};a[Symbol.iterator]=function(){return this};
return a}
function u(a){var b="undefined"!=typeof Symbol&&Symbol.iterator&&a[Symbol.iterator];return b?b.call(a):{next:aa(a)}}
var fa="function"==typeof Object.create?Object.create:function(a){function b(){}
b.prototype=a;return new b},ka;
if("function"==typeof Object.setPrototypeOf)ka=Object.setPrototypeOf;else{var la;a:{var ma={a:!0},na={};try{na.__proto__=ma;la=na.a;break a}catch(a){}la=!1}ka=la?function(a,b){a.__proto__=b;if(a.__proto__!==b)throw new TypeError(a+" is not extensible");return a}:null}var oa=ka;
function x(a,b){a.prototype=fa(b.prototype);a.prototype.constructor=a;if(oa)oa(a,b);else for(var c in b)if("prototype"!=c)if(Object.defineProperties){var d=Object.getOwnPropertyDescriptor(b,c);d&&Object.defineProperty(a,c,d)}else a[c]=b[c];a.A=b.prototype}
function pa(){this.o=!1;this.i=null;this.m=void 0;this.h=1;this.u=this.l=0;this.j=null}
function qa(a){if(a.o)throw new TypeError("Generator is already running");a.o=!0}
pa.prototype.s=function(a){this.m=a};
function ra(a,b){a.j={U:b,na:!0};a.h=a.l||a.u}
pa.prototype["return"]=function(a){this.j={"return":a};this.h=this.u};
function z(a,b,c){a.h=c;return{value:b}}
function va(a){a.l=0;var b=a.j.U;a.j=null;return b}
function wa(a){this.h=new pa;this.i=a}
function xa(a,b){qa(a.h);var c=a.h.i;if(c)return ya(a,"return"in c?c["return"]:function(d){return{value:d,done:!0}},b,a.h["return"]);
a.h["return"](b);return za(a)}
function ya(a,b,c,d){try{var e=b.call(a.h.i,c);if(!(e instanceof Object))throw new TypeError("Iterator result "+e+" is not an object");if(!e.done)return a.h.o=!1,e;var f=e.value}catch(g){return a.h.i=null,ra(a.h,g),za(a)}a.h.i=null;d.call(a.h,f);return za(a)}
function za(a){for(;a.h.h;)try{var b=a.i(a.h);if(b)return a.h.o=!1,{value:b.value,done:!1}}catch(c){a.h.m=void 0,ra(a.h,c)}a.h.o=!1;if(a.h.j){b=a.h.j;a.h.j=null;if(b.na)throw b.U;return{value:b["return"],done:!0}}return{value:void 0,done:!0}}
function Aa(a){this.next=function(b){qa(a.h);a.h.i?b=ya(a,a.h.i.next,b,a.h.s):(a.h.s(b),b=za(a));return b};
this["throw"]=function(b){qa(a.h);a.h.i?b=ya(a,a.h.i["throw"],b,a.h.s):(ra(a.h,b),b=za(a));return b};
this["return"]=function(b){return xa(a,b)};
this[Symbol.iterator]=function(){return this}}
function Ba(a,b){var c=new Aa(new wa(b));oa&&a.prototype&&oa(c,a.prototype);return c}
t("Reflect.setPrototypeOf",function(a){return a?a:oa?function(b,c){try{return oa(b,c),!0}catch(d){return!1}}:null});
t("Object.setPrototypeOf",function(a){return a||oa});
function A(a,b){return Object.prototype.hasOwnProperty.call(a,b)}
var Ca="function"==typeof Object.assign?Object.assign:function(a,b){for(var c=1;c<arguments.length;c++){var d=arguments[c];if(d)for(var e in d)A(d,e)&&(a[e]=d[e])}return a};
t("Object.assign",function(a){return a||Ca});
t("Promise",function(a){function b(g){this.h=0;this.j=void 0;this.i=[];this.s=!1;var h=this.l();try{g(h.resolve,h.reject)}catch(k){h.reject(k)}}
function c(){this.h=null}
function d(g){return g instanceof b?g:new b(function(h){h(g)})}
if(a)return a;c.prototype.i=function(g){if(null==this.h){this.h=[];var h=this;this.j(function(){h.m()})}this.h.push(g)};
var e=da.setTimeout;c.prototype.j=function(g){e(g,0)};
c.prototype.m=function(){for(;this.h&&this.h.length;){var g=this.h;this.h=[];for(var h=0;h<g.length;++h){var k=g[h];g[h]=null;try{k()}catch(l){this.l(l)}}}this.h=null};
c.prototype.l=function(g){this.j(function(){throw g;})};
b.prototype.l=function(){function g(l){return function(m){k||(k=!0,l.call(h,m))}}
var h=this,k=!1;return{resolve:g(this.ba),reject:g(this.m)}};
b.prototype.ba=function(g){if(g===this)this.m(new TypeError("A Promise cannot resolve to itself"));else if(g instanceof b)this.da(g);else{a:switch(typeof g){case "object":var h=null!=g;break a;case "function":h=!0;break a;default:h=!1}h?this.aa(g):this.o(g)}};
b.prototype.aa=function(g){var h=void 0;try{h=g.then}catch(k){this.m(k);return}"function"==typeof h?this.ea(h,g):this.o(g)};
b.prototype.m=function(g){this.u(2,g)};
b.prototype.o=function(g){this.u(1,g)};
b.prototype.u=function(g,h){if(0!=this.h)throw Error("Cannot settle("+g+", "+h+"): Promise already settled in state"+this.h);this.h=g;this.j=h;2===this.h&&this.ca();this.D()};
b.prototype.ca=function(){var g=this;e(function(){if(g.G()){var h=da.console;"undefined"!==typeof h&&h.error(g.j)}},1)};
b.prototype.G=function(){if(this.s)return!1;var g=da.CustomEvent,h=da.Event,k=da.dispatchEvent;if("undefined"===typeof k)return!0;"function"===typeof g?g=new g("unhandledrejection",{cancelable:!0}):"function"===typeof h?g=new h("unhandledrejection",{cancelable:!0}):(g=da.document.createEvent("CustomEvent"),g.initCustomEvent("unhandledrejection",!1,!0,g));g.promise=this;g.reason=this.j;return k(g)};
b.prototype.D=function(){if(null!=this.i){for(var g=0;g<this.i.length;++g)f.i(this.i[g]);this.i=null}};
var f=new c;b.prototype.da=function(g){var h=this.l();g.M(h.resolve,h.reject)};
b.prototype.ea=function(g,h){var k=this.l();try{g.call(h,k.resolve,k.reject)}catch(l){k.reject(l)}};
b.prototype.then=function(g,h){function k(p,q){return"function"==typeof p?function(v){try{l(p(v))}catch(w){m(w)}}:q}
var l,m,n=new b(function(p,q){l=p;m=q});
this.M(k(g,l),k(h,m));return n};
b.prototype["catch"]=function(g){return this.then(void 0,g)};
b.prototype.M=function(g,h){function k(){switch(l.h){case 1:g(l.j);break;case 2:h(l.j);break;default:throw Error("Unexpected state: "+l.h);}}
var l=this;null==this.i?f.i(k):this.i.push(k);this.s=!0};
b.resolve=d;b.reject=function(g){return new b(function(h,k){k(g)})};
b.race=function(g){return new b(function(h,k){for(var l=u(g),m=l.next();!m.done;m=l.next())d(m.value).M(h,k)})};
b.all=function(g){var h=u(g),k=h.next();return k.done?d([]):new b(function(l,m){function n(v){return function(w){p[v]=w;q--;0==q&&l(p)}}
var p=[],q=0;do p.push(void 0),q++,d(k.value).M(n(p.length-1),m),k=h.next();while(!k.done)})};
return b});
function Da(a,b,c){if(null==a)throw new TypeError("The 'this' value for String.prototype."+c+" must not be null or undefined");if(b instanceof RegExp)throw new TypeError("First argument to String.prototype."+c+" must not be a regular expression");return a+""}
t("String.prototype.endsWith",function(a){return a?a:function(b,c){var d=Da(this,b,"endsWith");b+="";void 0===c&&(c=d.length);for(var e=Math.max(0,Math.min(c|0,d.length)),f=b.length;0<f&&0<e;)if(d[--e]!=b[--f])return!1;return 0>=f}});
t("String.prototype.startsWith",function(a){return a?a:function(b,c){var d=Da(this,b,"startsWith");b+="";for(var e=d.length,f=b.length,g=Math.max(0,Math.min(c|0,d.length)),h=0;h<f&&g<e;)if(d[g++]!=b[h++])return!1;return h>=f}});
t("Object.is",function(a){return a?a:function(b,c){return b===c?0!==b||1/b===1/c:b!==b&&c!==c}});
t("Array.prototype.includes",function(a){return a?a:function(b,c){var d=this;d instanceof String&&(d=String(d));var e=d.length,f=c||0;for(0>f&&(f=Math.max(f+e,0));f<e;f++){var g=d[f];if(g===b||Object.is(g,b))return!0}return!1}});
t("String.prototype.includes",function(a){return a?a:function(b,c){return-1!==Da(this,b,"includes").indexOf(b,c||0)}});
t("Object.entries",function(a){return a?a:function(b){var c=[],d;for(d in b)A(b,d)&&c.push([d,b[d]]);return c}});
function Ga(a,b){a instanceof String&&(a+="");var c=0,d=!1,e={next:function(){if(!d&&c<a.length){var f=c++;return{value:b(f,a[f]),done:!1}}d=!0;return{done:!0,value:void 0}}};
e[Symbol.iterator]=function(){return e};
return e}
t("Array.prototype.entries",function(a){return a?a:function(){return Ga(this,function(b,c){return[b,c]})}});
t("Array.prototype.keys",function(a){return a?a:function(){return Ga(this,function(b){return b})}});
t("Array.prototype.values",function(a){return a?a:function(){return Ga(this,function(b,c){return c})}});
t("WeakMap",function(a){function b(k){this.h=(h+=Math.random()+1).toString();if(k){k=u(k);for(var l;!(l=k.next()).done;)l=l.value,this.set(l[0],l[1])}}
function c(){}
function d(k){var l=typeof k;return"object"===l&&null!==k||"function"===l}
function e(k){if(!A(k,g)){var l=new c;ba(k,g,{value:l})}}
function f(k){var l=Object[k];l&&(Object[k]=function(m){if(m instanceof c)return m;Object.isExtensible(m)&&e(m);return l(m)})}
if(function(){if(!a||!Object.seal)return!1;try{var k=Object.seal({}),l=Object.seal({}),m=new a([[k,2],[l,3]]);if(2!=m.get(k)||3!=m.get(l))return!1;m["delete"](k);m.set(l,4);return!m.has(k)&&4==m.get(l)}catch(n){return!1}}())return a;
var g="$jscomp_hidden_"+Math.random();f("freeze");f("preventExtensions");f("seal");var h=0;b.prototype.set=function(k,l){if(!d(k))throw Error("Invalid WeakMap key");e(k);if(!A(k,g))throw Error("WeakMap key fail: "+k);k[g][this.h]=l;return this};
b.prototype.get=function(k){return d(k)&&A(k,g)?k[g][this.h]:void 0};
b.prototype.has=function(k){return d(k)&&A(k,g)&&A(k[g],this.h)};
b.prototype["delete"]=function(k){return d(k)&&A(k,g)&&A(k[g],this.h)?delete k[g][this.h]:!1};
return b});
t("Number.isNaN",function(a){return a?a:function(b){return"number"===typeof b&&isNaN(b)}});
t("Map",function(a){function b(){var h={};return h.previous=h.next=h.head=h}
function c(h,k){var l=h.h;return ea(function(){if(l){for(;l.head!=h.h;)l=l.previous;for(;l.next!=l.head;)return l=l.next,{done:!1,value:k(l)};l=null}return{done:!0,value:void 0}})}
function d(h,k){var l=k&&typeof k;"object"==l||"function"==l?f.has(k)?l=f.get(k):(l=""+ ++g,f.set(k,l)):l="p_"+k;var m=h.i[l];if(m&&A(h.i,l))for(var n=0;n<m.length;n++){var p=m[n];if(k!==k&&p.key!==p.key||k===p.key)return{id:l,list:m,index:n,v:p}}return{id:l,list:m,index:-1,v:void 0}}
function e(h){this.i={};this.h=b();this.size=0;if(h){h=u(h);for(var k;!(k=h.next()).done;)k=k.value,this.set(k[0],k[1])}}
if(function(){if(!a||"function"!=typeof a||!a.prototype.entries||"function"!=typeof Object.seal)return!1;try{var h=Object.seal({x:4}),k=new a(u([[h,"s"]]));if("s"!=k.get(h)||1!=k.size||k.get({x:4})||k.set({x:4},"t")!=k||2!=k.size)return!1;var l=k.entries(),m=l.next();if(m.done||m.value[0]!=h||"s"!=m.value[1])return!1;m=l.next();return m.done||4!=m.value[0].x||"t"!=m.value[1]||!l.next().done?!1:!0}catch(n){return!1}}())return a;
var f=new WeakMap;e.prototype.set=function(h,k){h=0===h?0:h;var l=d(this,h);l.list||(l.list=this.i[l.id]=[]);l.v?l.v.value=k:(l.v={next:this.h,previous:this.h.previous,head:this.h,key:h,value:k},l.list.push(l.v),this.h.previous.next=l.v,this.h.previous=l.v,this.size++);return this};
e.prototype["delete"]=function(h){h=d(this,h);return h.v&&h.list?(h.list.splice(h.index,1),h.list.length||delete this.i[h.id],h.v.previous.next=h.v.next,h.v.next.previous=h.v.previous,h.v.head=null,this.size--,!0):!1};
e.prototype.clear=function(){this.i={};this.h=this.h.previous=b();this.size=0};
e.prototype.has=function(h){return!!d(this,h).v};
e.prototype.get=function(h){return(h=d(this,h).v)&&h.value};
e.prototype.entries=function(){return c(this,function(h){return[h.key,h.value]})};
e.prototype.keys=function(){return c(this,function(h){return h.key})};
e.prototype.values=function(){return c(this,function(h){return h.value})};
e.prototype.forEach=function(h,k){for(var l=this.entries(),m;!(m=l.next()).done;)m=m.value,h.call(k,m[1],m[0],this)};
e.prototype[Symbol.iterator]=e.prototype.entries;var g=0;return e});
t("Set",function(a){function b(c){this.h=new Map;if(c){c=u(c);for(var d;!(d=c.next()).done;)this.add(d.value)}this.size=this.h.size}
if(function(){if(!a||"function"!=typeof a||!a.prototype.entries||"function"!=typeof Object.seal)return!1;try{var c=Object.seal({x:4}),d=new a(u([c]));if(!d.has(c)||1!=d.size||d.add(c)!=d||1!=d.size||d.add({x:4})!=d||2!=d.size)return!1;var e=d.entries(),f=e.next();if(f.done||f.value[0]!=c||f.value[1]!=c)return!1;f=e.next();return f.done||f.value[0]==c||4!=f.value[0].x||f.value[1]!=f.value[0]?!1:e.next().done}catch(g){return!1}}())return a;
b.prototype.add=function(c){c=0===c?0:c;this.h.set(c,c);this.size=this.h.size;return this};
b.prototype["delete"]=function(c){c=this.h["delete"](c);this.size=this.h.size;return c};
b.prototype.clear=function(){this.h.clear();this.size=0};
b.prototype.has=function(c){return this.h.has(c)};
b.prototype.entries=function(){return this.h.entries()};
b.prototype.values=function(){return this.h.values()};
b.prototype.keys=b.prototype.values;b.prototype[Symbol.iterator]=b.prototype.values;b.prototype.forEach=function(c,d){var e=this;this.h.forEach(function(f){return c.call(d,f,f,e)})};
return b});
var B=this||self;function C(a,b){for(var c=a.split("."),d=b||B,e=0;e<c.length;e++)if(d=d[c[e]],null==d)return null;return d}
function Ha(){}
function Ja(a){var b=typeof a;b="object"!=b?b:a?Array.isArray(a)?"array":b:"null";return"array"==b||"object"==b&&"number"==typeof a.length}
function D(a){var b=typeof a;return"object"==b&&null!=a||"function"==b}
function Ka(a){return Object.prototype.hasOwnProperty.call(a,La)&&a[La]||(a[La]=++Ma)}
var La="closure_uid_"+(1E9*Math.random()>>>0),Ma=0;function Na(a,b,c){return a.call.apply(a.bind,arguments)}
function Oa(a,b,c){if(!a)throw Error();if(2<arguments.length){var d=Array.prototype.slice.call(arguments,2);return function(){var e=Array.prototype.slice.call(arguments);Array.prototype.unshift.apply(e,d);return a.apply(b,e)}}return function(){return a.apply(b,arguments)}}
function Pa(a,b,c){Function.prototype.bind&&-1!=Function.prototype.bind.toString().indexOf("native code")?Pa=Na:Pa=Oa;return Pa.apply(null,arguments)}
function E(a,b){var c=a.split("."),d=B;c[0]in d||"undefined"==typeof d.execScript||d.execScript("var "+c[0]);for(var e;c.length&&(e=c.shift());)c.length||void 0===b?d[e]&&d[e]!==Object.prototype[e]?d=d[e]:d=d[e]={}:d[e]=b}
function F(a,b){function c(){}
c.prototype=b.prototype;a.A=b.prototype;a.prototype=new c;a.prototype.constructor=a;a.Ha=function(d,e,f){for(var g=Array(arguments.length-2),h=2;h<arguments.length;h++)g[h-2]=arguments[h];return b.prototype[e].apply(d,g)}}
function Qa(a){return a}
;function Ra(a,b){var c=void 0;return new (c||(c=Promise))(function(d,e){function f(k){try{h(b.next(k))}catch(l){e(l)}}
function g(k){try{h(b["throw"](k))}catch(l){e(l)}}
function h(k){k.done?d(k.value):(new c(function(l){l(k.value)})).then(f,g)}
h((b=b.apply(a,void 0)).next())})}
;function Sa(a){if(Error.captureStackTrace)Error.captureStackTrace(this,Sa);else{var b=Error().stack;b&&(this.stack=b)}a&&(this.message=String(a))}
F(Sa,Error);Sa.prototype.name="CustomError";var Ta=Array.prototype.indexOf?function(a,b){return Array.prototype.indexOf.call(a,b,void 0)}:function(a,b){if("string"===typeof a)return"string"!==typeof b||1!=b.length?-1:a.indexOf(b,0);
for(var c=0;c<a.length;c++)if(c in a&&a[c]===b)return c;return-1},H=Array.prototype.forEach?function(a,b,c){Array.prototype.forEach.call(a,b,c)}:function(a,b,c){for(var d=a.length,e="string"===typeof a?a.split(""):a,f=0;f<d;f++)f in e&&b.call(c,e[f],f,a)},Ya=Array.prototype.reduce?function(a,b,c){return Array.prototype.reduce.call(a,b,c)}:function(a,b,c){var d=c;
H(a,function(e,f){d=b.call(void 0,d,e,f,a)});
return d};
function Za(a,b){a:{var c=a.length;for(var d="string"===typeof a?a.split(""):a,e=0;e<c;e++)if(e in d&&b.call(void 0,d[e],e,a)){c=e;break a}c=-1}return 0>c?null:"string"===typeof a?a.charAt(c):a[c]}
function $a(a,b){var c=Ta(a,b),d;(d=0<=c)&&Array.prototype.splice.call(a,c,1);return d}
function ab(a){return Array.prototype.concat.apply([],arguments)}
function bb(a){var b=a.length;if(0<b){for(var c=Array(b),d=0;d<b;d++)c[d]=a[d];return c}return[]}
function cb(a,b){for(var c=1;c<arguments.length;c++){var d=arguments[c];if(Ja(d)){var e=a.length||0,f=d.length||0;a.length=e+f;for(var g=0;g<f;g++)a[e+g]=d[g]}else a.push(d)}}
;function db(a){var b=!1,c;return function(){b||(c=a(),b=!0);return c}}
;function eb(a,b){for(var c in a)b.call(void 0,a[c],c,a)}
function fb(a){var b=gb,c;for(c in b)if(a.call(void 0,b[c],c,b))return c}
function hb(a,b){for(var c in a)if(!(c in b)||a[c]!==b[c])return!1;for(var d in b)if(!(d in a))return!1;return!0}
function ib(a){if(!a||"object"!==typeof a)return a;if("function"===typeof a.clone)return a.clone();var b=Array.isArray(a)?[]:"function"!==typeof ArrayBuffer||"function"!==typeof ArrayBuffer.isView||!ArrayBuffer.isView(a)||a instanceof DataView?{}:new a.constructor(a.length),c;for(c in a)b[c]=ib(a[c]);return b}
var jb="constructor hasOwnProperty isPrototypeOf propertyIsEnumerable toLocaleString toString valueOf".split(" ");function kb(a,b){for(var c,d,e=1;e<arguments.length;e++){d=arguments[e];for(c in d)a[c]=d[c];for(var f=0;f<jb.length;f++)c=jb[f],Object.prototype.hasOwnProperty.call(d,c)&&(a[c]=d[c])}}
;var lb;var mb=String.prototype.trim?function(a){return a.trim()}:function(a){return/^[\s\xa0]*([\s\S]*?)[\s\xa0]*$/.exec(a)[1]},nb=/&/g,ob=/</g,pb=/>/g,qb=/"/g,rb=/'/g,yb=/\x00/g,zb=/[\x00&<>"']/;
function Ab(a,b){return a<b?-1:a>b?1:0}
;var Bb;a:{var Cb=B.navigator;if(Cb){var Db=Cb.userAgent;if(Db){Bb=Db;break a}}Bb=""}function I(a){return-1!=Bb.indexOf(a)}
;function Eb(a){this.h=Fb===Fb?a:""}
Eb.prototype.toString=function(){return this.h.toString()};
var Fb={};function Gb(){return I("iPhone")&&!I("iPod")&&!I("iPad")}
;function Hb(a){Hb[" "](a);return a}
Hb[" "]=Ha;var Ib=I("Opera"),Jb=I("Trident")||I("MSIE"),Kb=I("Edge"),Lb=I("Gecko")&&!(-1!=Bb.toLowerCase().indexOf("webkit")&&!I("Edge"))&&!(I("Trident")||I("MSIE"))&&!I("Edge"),Mb=-1!=Bb.toLowerCase().indexOf("webkit")&&!I("Edge");function Nb(){var a=B.document;return a?a.documentMode:void 0}
var Ob;a:{var Pb="",Qb=function(){var a=Bb;if(Lb)return/rv:([^\);]+)(\)|;)/.exec(a);if(Kb)return/Edge\/([\d\.]+)/.exec(a);if(Jb)return/\b(?:MSIE|rv)[: ]([^\);]+)(\)|;)/.exec(a);if(Mb)return/WebKit\/(\S+)/.exec(a);if(Ib)return/(?:Version)[ \/]?(\S+)/.exec(a)}();
Qb&&(Pb=Qb?Qb[1]:"");if(Jb){var Rb=Nb();if(null!=Rb&&Rb>parseFloat(Pb)){Ob=String(Rb);break a}}Ob=Pb}var Sb=Ob,Tb={},Ub;if(B.document&&Jb){var Vb=Nb();Ub=Vb?Vb:parseInt(Sb,10)||void 0}else Ub=void 0;var Wb=Ub;var Xb=I("Firefox")||I("FxiOS"),Yb=Gb()||I("iPod"),Zb=I("iPad"),$b=I("Safari")&&!((I("Chrome")||I("CriOS"))&&!I("Edge")||I("Coast")||I("Opera")||I("Edge")||I("Edg/")||I("OPR")||I("Firefox")||I("FxiOS")||I("Silk")||I("Android"))&&!(Gb()||I("iPad")||I("iPod"));var ac={},bc=null;var J=window;function hc(a,b){this.width=a;this.height=b}
r=hc.prototype;r.clone=function(){return new hc(this.width,this.height)};
r.aspectRatio=function(){return this.width/this.height};
r.isEmpty=function(){return!(this.width*this.height)};
r.ceil=function(){this.width=Math.ceil(this.width);this.height=Math.ceil(this.height);return this};
r.floor=function(){this.width=Math.floor(this.width);this.height=Math.floor(this.height);return this};
r.round=function(){this.width=Math.round(this.width);this.height=Math.round(this.height);return this};function ic(){var a=document;var b="IFRAME";"application/xhtml+xml"===a.contentType&&(b=b.toLowerCase());return a.createElement(b)}
function jc(a,b){for(var c=0;a;){if(b(a))return a;a=a.parentNode;c++}return null}
;var kc=/^(?:([^:/?#.]+):)?(?:\/\/(?:([^\\/?#]*)@)?([^\\/?#]*?)(?::([0-9]+))?(?=[\\/?#]|$))?([^?#]+)?(?:\?([^#]*))?(?:#([\s\S]*))?$/;function lc(a){return a?decodeURI(a):a}
function mc(a){return lc(a.match(kc)[3]||null)}
function nc(a){var b=a.match(kc);a=b[1];var c=b[2],d=b[3];b=b[4];var e="";a&&(e+=a+":");d&&(e+="//",c&&(e+=c+"@"),e+=d,b&&(e+=":"+b));return e}
function oc(a,b,c){if(Array.isArray(b))for(var d=0;d<b.length;d++)oc(a,String(b[d]),c);else null!=b&&c.push(a+(""===b?"":"="+encodeURIComponent(String(b))))}
function pc(a){var b=[],c;for(c in a)oc(c,a[c],b);return b.join("&")}
var qc=/#|$/;function rc(a){var b=sc;if(b)for(var c in b)Object.prototype.hasOwnProperty.call(b,c)&&a.call(void 0,b[c],c,b)}
function tc(){var a=[];rc(function(b){a.push(b)});
return a}
var sc={ta:"allow-forms",va:"allow-modals",wa:"allow-orientation-lock",xa:"allow-pointer-lock",ya:"allow-popups",za:"allow-popups-to-escape-sandbox",Aa:"allow-presentation",Ba:"allow-same-origin",Ca:"allow-scripts",Da:"allow-top-navigation",Ea:"allow-top-navigation-by-user-activation"},uc=db(function(){return tc()});
function vc(){var a=ic(),b={};H(uc(),function(c){a.sandbox&&a.sandbox.supports&&a.sandbox.supports(c)&&(b[c]=!0)});
return b}
;function wc(){this.j=this.j;this.m=this.m}
wc.prototype.j=!1;wc.prototype.dispose=function(){this.j||(this.j=!0,this.J())};
wc.prototype.J=function(){if(this.m)for(;this.m.length;)this.m.shift()()};var xc={};function yc(a){if(a!==xc)throw Error("Bad secret");}
;function zc(){var a="undefined"!==typeof window?window.trustedTypes:void 0;return null!==a&&void 0!==a?a:null}
;var Ac;function Bc(){}
function Cc(a,b){yc(b);this.h=a}
x(Cc,Bc);Cc.prototype.toString=function(){return this.h.toString()};
var Dc=null===(Ac=zc())||void 0===Ac?void 0:Ac.emptyHTML;new Cc(null!==Dc&&void 0!==Dc?Dc:"",xc);var Ec;function Fc(){}
function Gc(a,b){yc(b);this.h=a}
x(Gc,Fc);Gc.prototype.toString=function(){return this.h.toString()};
var Hc=null===(Ec=zc())||void 0===Ec?void 0:Ec.emptyScript;new Gc(null!==Hc&&void 0!==Hc?Hc:"",xc);function Ic(){}
function Jc(a,b){yc(b);this.h=a}
x(Jc,Ic);Jc.prototype.toString=function(){return this.h};
new Jc("about:blank",xc);new Jc("about:invalid#zTSz",xc);function Kc(a){var b=C("window.location.href");null==a&&(a='Unknown Error of type "null/undefined"');if("string"===typeof a)return{message:a,name:"Unknown error",lineNumber:"Not available",fileName:b,stack:"Not available"};var c=!1;try{var d=a.lineNumber||a.line||"Not available"}catch(g){d="Not available",c=!0}try{var e=a.fileName||a.filename||a.sourceURL||B.$googDebugFname||b}catch(g){e="Not available",c=!0}b=Lc(a);if(!(!c&&a.lineNumber&&a.fileName&&a.stack&&a.message&&a.name)){c=a.message;if(null==
c){if(a.constructor&&a.constructor instanceof Function){if(a.constructor.name)c=a.constructor.name;else if(c=a.constructor,Mc[c])c=Mc[c];else{c=String(c);if(!Mc[c]){var f=/function\s+([^\(]+)/m.exec(c);Mc[c]=f?f[1]:"[Anonymous]"}c=Mc[c]}c='Unknown Error of type "'+c+'"'}else c="Unknown Error of unknown type";"function"===typeof a.toString&&Object.prototype.toString!==a.toString&&(c+=": "+a.toString())}return{message:c,name:a.name||"UnknownError",lineNumber:d,fileName:e,stack:b||"Not available"}}a.stack=
b;return a}
function Lc(a,b){b||(b={});b[Nc(a)]=!0;var c=a.stack||"",d=a.Ja;d&&!b[Nc(d)]&&(c+="\nCaused by: ",d.stack&&0==d.stack.indexOf(d.toString())||(c+="string"===typeof d?d:d.message+"\n"),c+=Lc(d,b));return c}
function Nc(a){var b="";"function"===typeof a.toString&&(b=""+a);return b+a.stack}
var Mc={};function Oc(a){this.h=a||{cookie:""}}
r=Oc.prototype;r.isEnabled=function(){return navigator.cookieEnabled};
r.set=function(a,b,c){var d=!1;if("object"===typeof c){var e=c.Pa;d=c.secure||!1;var f=c.domain||void 0;var g=c.path||void 0;var h=c.V}if(/[;=\s]/.test(a))throw Error('Invalid cookie name "'+a+'"');if(/[;\r\n]/.test(b))throw Error('Invalid cookie value "'+b+'"');void 0===h&&(h=-1);this.h.cookie=a+"="+b+(f?";domain="+f:"")+(g?";path="+g:"")+(0>h?"":0==h?";expires="+(new Date(1970,1,1)).toUTCString():";expires="+(new Date(Date.now()+1E3*h)).toUTCString())+(d?";secure":"")+(null!=e?";samesite="+e:"")};
r.get=function(a,b){for(var c=a+"=",d=(this.h.cookie||"").split(";"),e=0,f;e<d.length;e++){f=mb(d[e]);if(0==f.lastIndexOf(c,0))return f.substr(c.length);if(f==a)return""}return b};
r.remove=function(a,b,c){var d=void 0!==this.get(a);this.set(a,"",{V:0,path:b,domain:c});return d};
r.isEmpty=function(){return!this.h.cookie};
r.clear=function(){for(var a=(this.h.cookie||"").split(";"),b=[],c=[],d,e,f=0;f<a.length;f++)e=mb(a[f]),d=e.indexOf("="),-1==d?(b.push(""),c.push(e)):(b.push(e.substring(0,d)),c.push(e.substring(d+1)));for(a=b.length-1;0<=a;a--)this.remove(b[a])};
var Pc=new Oc("undefined"==typeof document?null:document);var Qc=(new Date).getTime();function Rc(a){if(!a)return"";a=a.split("#")[0].split("?")[0];a=a.toLowerCase();0==a.indexOf("//")&&(a=window.location.protocol+a);/^[\w\-]*:\/\//.test(a)||(a=window.location.href);var b=a.substring(a.indexOf("://")+3),c=b.indexOf("/");-1!=c&&(b=b.substring(0,c));c=a.substring(0,a.indexOf("://"));if(!c)throw Error("URI is missing protocol: "+a);if("http"!==c&&"https"!==c&&"chrome-extension"!==c&&"moz-extension"!==c&&"file"!==c&&"android-app"!==c&&"chrome-search"!==c&&"chrome-untrusted"!==c&&"chrome"!==
c&&"app"!==c&&"devtools"!==c)throw Error("Invalid URI scheme in origin: "+c);a="";var d=b.indexOf(":");if(-1!=d){var e=b.substring(d+1);b=b.substring(0,d);if("http"===c&&"80"!==e||"https"===c&&"443"!==e)a=":"+e}return c+"://"+b+a}
;function Sc(){function a(){e[0]=1732584193;e[1]=4023233417;e[2]=2562383102;e[3]=271733878;e[4]=3285377520;m=l=0}
function b(n){for(var p=g,q=0;64>q;q+=4)p[q/4]=n[q]<<24|n[q+1]<<16|n[q+2]<<8|n[q+3];for(q=16;80>q;q++)n=p[q-3]^p[q-8]^p[q-14]^p[q-16],p[q]=(n<<1|n>>>31)&4294967295;n=e[0];var v=e[1],w=e[2],y=e[3],G=e[4];for(q=0;80>q;q++){if(40>q)if(20>q){var Q=y^v&(w^y);var Ia=1518500249}else Q=v^w^y,Ia=1859775393;else 60>q?(Q=v&w|y&(v|w),Ia=2400959708):(Q=v^w^y,Ia=3395469782);Q=((n<<5|n>>>27)&4294967295)+Q+G+Ia+p[q]&4294967295;G=y;y=w;w=(v<<30|v>>>2)&4294967295;v=n;n=Q}e[0]=e[0]+n&4294967295;e[1]=e[1]+v&4294967295;
e[2]=e[2]+w&4294967295;e[3]=e[3]+y&4294967295;e[4]=e[4]+G&4294967295}
function c(n,p){if("string"===typeof n){n=unescape(encodeURIComponent(n));for(var q=[],v=0,w=n.length;v<w;++v)q.push(n.charCodeAt(v));n=q}p||(p=n.length);q=0;if(0==l)for(;q+64<p;)b(n.slice(q,q+64)),q+=64,m+=64;for(;q<p;)if(f[l++]=n[q++],m++,64==l)for(l=0,b(f);q+64<p;)b(n.slice(q,q+64)),q+=64,m+=64}
function d(){var n=[],p=8*m;56>l?c(h,56-l):c(h,64-(l-56));for(var q=63;56<=q;q--)f[q]=p&255,p>>>=8;b(f);for(q=p=0;5>q;q++)for(var v=24;0<=v;v-=8)n[p++]=e[q]>>v&255;return n}
for(var e=[],f=[],g=[],h=[128],k=1;64>k;++k)h[k]=0;var l,m;a();return{reset:a,update:c,digest:d,fa:function(){for(var n=d(),p="",q=0;q<n.length;q++)p+="0123456789ABCDEF".charAt(Math.floor(n[q]/16))+"0123456789ABCDEF".charAt(n[q]%16);return p}}}
;function Tc(a,b,c){var d=[],e=[];if(1==(Array.isArray(c)?2:1))return e=[b,a],H(d,function(h){e.push(h)}),Uc(e.join(" "));
var f=[],g=[];H(c,function(h){g.push(h.key);f.push(h.value)});
c=Math.floor((new Date).getTime()/1E3);e=0==f.length?[c,b,a]:[f.join(":"),c,b,a];H(d,function(h){e.push(h)});
a=Uc(e.join(" "));a=[c,a];0==g.length||a.push(g.join(""));return a.join("_")}
function Uc(a){var b=Sc();b.update(a);return b.fa().toLowerCase()}
;function Vc(a){var b=Rc(String(B.location.href)),c;(c=B.__SAPISID||B.__APISID||B.__OVERRIDE_SID)?c=!0:(c=new Oc(document),c=c.get("SAPISID")||c.get("APISID")||c.get("__Secure-3PAPISID")||c.get("SID"),c=!!c);if(c&&(c=(b=0==b.indexOf("https:")||0==b.indexOf("chrome-extension:")||0==b.indexOf("moz-extension:"))?B.__SAPISID:B.__APISID,c||(c=new Oc(document),c=c.get(b?"SAPISID":"APISID")||c.get("__Secure-3PAPISID")),c)){b=b?"SAPISIDHASH":"APISIDHASH";var d=String(B.location.href);return d&&c&&b?[b,Tc(Rc(d),
c,a||null)].join(" "):null}return null}
;function Wc(){this.i=[];this.h=-1}
Wc.prototype.set=function(a,b){b=void 0===b?!0:b;0<=a&&52>a&&0===a%1&&this.i[a]!=b&&(this.i[a]=b,this.h=-1)};
Wc.prototype.get=function(a){return!!this.i[a]};
function Xc(a){-1==a.h&&(a.h=Ya(a.i,function(b,c,d){return c?b+Math.pow(2,d):b},0));
return a.h}
;function Yc(a,b){this.j=a;this.l=b;this.i=0;this.h=null}
Yc.prototype.get=function(){if(0<this.i){this.i--;var a=this.h;this.h=a.next;a.next=null}else a=this.j();return a};
function Zc(a,b){a.l(b);100>a.i&&(a.i++,b.next=a.h,a.h=b)}
;var $c;function ad(){var a=B.MessageChannel;"undefined"===typeof a&&"undefined"!==typeof window&&window.postMessage&&window.addEventListener&&!I("Presto")&&(a=function(){var e=ic();e.style.display="none";document.documentElement.appendChild(e);var f=e.contentWindow;e=f.document;e.open();e.close();var g="callImmediate"+Math.random(),h="file:"==f.location.protocol?"*":f.location.protocol+"//"+f.location.host;e=Pa(function(k){if(("*"==h||k.origin==h)&&k.data==g)this.port1.onmessage()},this);
f.addEventListener("message",e,!1);this.port1={};this.port2={postMessage:function(){f.postMessage(g,h)}}});
if("undefined"!==typeof a&&!I("Trident")&&!I("MSIE")){var b=new a,c={},d=c;b.port1.onmessage=function(){if(void 0!==c.next){c=c.next;var e=c.S;c.S=null;e()}};
return function(e){d.next={S:e};d=d.next;b.port2.postMessage(0)}}return function(e){B.setTimeout(e,0)}}
;function bd(a){B.setTimeout(function(){throw a;},0)}
;function cd(){this.i=this.h=null}
cd.prototype.add=function(a,b){var c=dd.get();c.set(a,b);this.i?this.i.next=c:this.h=c;this.i=c};
cd.prototype.remove=function(){var a=null;this.h&&(a=this.h,this.h=this.h.next,this.h||(this.i=null),a.next=null);return a};
var dd=new Yc(function(){return new ed},function(a){return a.reset()});
function ed(){this.next=this.scope=this.h=null}
ed.prototype.set=function(a,b){this.h=a;this.scope=b;this.next=null};
ed.prototype.reset=function(){this.next=this.scope=this.h=null};function fd(a,b){gd||hd();id||(gd(),id=!0);jd.add(a,b)}
var gd;function hd(){if(B.Promise&&B.Promise.resolve){var a=B.Promise.resolve(void 0);gd=function(){a.then(kd)}}else gd=function(){var b=kd;
"function"!==typeof B.setImmediate||B.Window&&B.Window.prototype&&!I("Edge")&&B.Window.prototype.setImmediate==B.setImmediate?($c||($c=ad()),$c(b)):B.setImmediate(b)}}
var id=!1,jd=new cd;function kd(){for(var a;a=jd.remove();){try{a.h.call(a.scope)}catch(b){bd(b)}Zc(dd,a)}id=!1}
;function ld(){this.i=-1}
;function md(){this.i=64;this.h=[];this.o=[];this.s=[];this.l=[];this.l[0]=128;for(var a=1;a<this.i;++a)this.l[a]=0;this.m=this.j=0;this.reset()}
F(md,ld);md.prototype.reset=function(){this.h[0]=1732584193;this.h[1]=4023233417;this.h[2]=2562383102;this.h[3]=271733878;this.h[4]=3285377520;this.m=this.j=0};
function nd(a,b,c){c||(c=0);var d=a.s;if("string"===typeof b)for(var e=0;16>e;e++)d[e]=b.charCodeAt(c)<<24|b.charCodeAt(c+1)<<16|b.charCodeAt(c+2)<<8|b.charCodeAt(c+3),c+=4;else for(e=0;16>e;e++)d[e]=b[c]<<24|b[c+1]<<16|b[c+2]<<8|b[c+3],c+=4;for(e=16;80>e;e++){var f=d[e-3]^d[e-8]^d[e-14]^d[e-16];d[e]=(f<<1|f>>>31)&4294967295}b=a.h[0];c=a.h[1];var g=a.h[2],h=a.h[3],k=a.h[4];for(e=0;80>e;e++){if(40>e)if(20>e){f=h^c&(g^h);var l=1518500249}else f=c^g^h,l=1859775393;else 60>e?(f=c&g|h&(c|g),l=2400959708):
(f=c^g^h,l=3395469782);f=(b<<5|b>>>27)+f+k+l+d[e]&4294967295;k=h;h=g;g=(c<<30|c>>>2)&4294967295;c=b;b=f}a.h[0]=a.h[0]+b&4294967295;a.h[1]=a.h[1]+c&4294967295;a.h[2]=a.h[2]+g&4294967295;a.h[3]=a.h[3]+h&4294967295;a.h[4]=a.h[4]+k&4294967295}
md.prototype.update=function(a,b){if(null!=a){void 0===b&&(b=a.length);for(var c=b-this.i,d=0,e=this.o,f=this.j;d<b;){if(0==f)for(;d<=c;)nd(this,a,d),d+=this.i;if("string"===typeof a)for(;d<b;){if(e[f]=a.charCodeAt(d),++f,++d,f==this.i){nd(this,e);f=0;break}}else for(;d<b;)if(e[f]=a[d],++f,++d,f==this.i){nd(this,e);f=0;break}}this.j=f;this.m+=b}};
md.prototype.digest=function(){var a=[],b=8*this.m;56>this.j?this.update(this.l,56-this.j):this.update(this.l,this.i-(this.j-56));for(var c=this.i-1;56<=c;c--)this.o[c]=b&255,b/=256;nd(this,this.o);for(c=b=0;5>c;c++)for(var d=24;0<=d;d-=8)a[b]=this.h[c]>>d&255,++b;return a};var od="StopIteration"in B?B.StopIteration:{message:"StopIteration",stack:""};function pd(){}
pd.prototype.next=function(){throw od;};
pd.prototype.B=function(){return this};
function qd(a){if(a instanceof pd)return a;if("function"==typeof a.B)return a.B(!1);if(Ja(a)){var b=0,c=new pd;c.next=function(){for(;;){if(b>=a.length)throw od;if(b in a)return a[b++];b++}};
return c}throw Error("Not implemented");}
function rd(a,b){if(Ja(a))try{H(a,b,void 0)}catch(c){if(c!==od)throw c;}else{a=qd(a);try{for(;;)b.call(void 0,a.next(),void 0,a)}catch(c){if(c!==od)throw c;}}}
function sd(a){if(Ja(a))return bb(a);a=qd(a);var b=[];rd(a,function(c){b.push(c)});
return b}
;function td(a,b){this.j={};this.h=[];this.l=this.i=0;var c=arguments.length;if(1<c){if(c%2)throw Error("Uneven number of arguments");for(var d=0;d<c;d+=2)this.set(arguments[d],arguments[d+1])}else if(a)if(a instanceof td)for(c=ud(a),d=0;d<c.length;d++)this.set(c[d],a.get(c[d]));else for(d in a)this.set(d,a[d])}
function ud(a){vd(a);return a.h.concat()}
r=td.prototype;r.equals=function(a,b){if(this===a)return!0;if(this.i!=a.i)return!1;var c=b||wd;vd(this);for(var d,e=0;d=this.h[e];e++)if(!c(this.get(d),a.get(d)))return!1;return!0};
function wd(a,b){return a===b}
r.isEmpty=function(){return 0==this.i};
r.clear=function(){this.j={};this.l=this.i=this.h.length=0};
r.remove=function(a){return Object.prototype.hasOwnProperty.call(this.j,a)?(delete this.j[a],this.i--,this.l++,this.h.length>2*this.i&&vd(this),!0):!1};
function vd(a){if(a.i!=a.h.length){for(var b=0,c=0;b<a.h.length;){var d=a.h[b];Object.prototype.hasOwnProperty.call(a.j,d)&&(a.h[c++]=d);b++}a.h.length=c}if(a.i!=a.h.length){var e={};for(c=b=0;b<a.h.length;)d=a.h[b],Object.prototype.hasOwnProperty.call(e,d)||(a.h[c++]=d,e[d]=1),b++;a.h.length=c}}
r.get=function(a,b){return Object.prototype.hasOwnProperty.call(this.j,a)?this.j[a]:b};
r.set=function(a,b){Object.prototype.hasOwnProperty.call(this.j,a)||(this.i++,this.h.push(a),this.l++);this.j[a]=b};
r.forEach=function(a,b){for(var c=ud(this),d=0;d<c.length;d++){var e=c[d],f=this.get(e);a.call(b,f,e,this)}};
r.clone=function(){return new td(this)};
r.B=function(a){vd(this);var b=0,c=this.l,d=this,e=new pd;e.next=function(){if(c!=d.l)throw Error("The map has changed since the iterator was created");if(b>=d.h.length)throw od;var f=d.h[b++];return a?f:d.j[f]};
return e};var xd=!Jb||9<=Number(Wb),Od;
if(Od=Jb){var Pd;if(Object.prototype.hasOwnProperty.call(Tb,"9"))Pd=Tb["9"];else{for(var Qd=0,Rd=mb(String(Sb)).split("."),Sd=mb("9").split("."),Td=Math.max(Rd.length,Sd.length),Ud=0;0==Qd&&Ud<Td;Ud++){var Vd=Rd[Ud]||"",Wd=Sd[Ud]||"";do{var Xd=/(\d*)(\D*)(.*)/.exec(Vd)||["","","",""],Yd=/(\d*)(\D*)(.*)/.exec(Wd)||["","","",""];if(0==Xd[0].length&&0==Yd[0].length)break;Qd=Ab(0==Xd[1].length?0:parseInt(Xd[1],10),0==Yd[1].length?0:parseInt(Yd[1],10))||Ab(0==Xd[2].length,0==Yd[2].length)||Ab(Xd[2],Yd[2]);
Vd=Xd[3];Wd=Yd[3]}while(0==Qd)}Pd=Tb["9"]=0<=Qd}Od=!Pd}var Zd=Od,$d=function(){if(!B.addEventListener||!Object.defineProperty)return!1;var a=!1,b=Object.defineProperty({},"passive",{get:function(){a=!0}});
try{B.addEventListener("test",Ha,b),B.removeEventListener("test",Ha,b)}catch(c){}return a}();function ae(a,b){this.type=a;this.h=this.target=b;this.defaultPrevented=this.i=!1}
ae.prototype.stopPropagation=function(){this.i=!0};
ae.prototype.preventDefault=function(){this.defaultPrevented=!0};function be(a,b){ae.call(this,a?a.type:"");this.relatedTarget=this.h=this.target=null;this.button=this.screenY=this.screenX=this.clientY=this.clientX=0;this.key="";this.charCode=this.keyCode=0;this.metaKey=this.shiftKey=this.altKey=this.ctrlKey=!1;this.state=null;this.pointerId=0;this.pointerType="";this.j=null;a&&this.init(a,b)}
F(be,ae);var ce={2:"touch",3:"pen",4:"mouse"};
be.prototype.init=function(a,b){var c=this.type=a.type,d=a.changedTouches&&a.changedTouches.length?a.changedTouches[0]:null;this.target=a.target||a.srcElement;this.h=b;var e=a.relatedTarget;if(e){if(Lb){a:{try{Hb(e.nodeName);var f=!0;break a}catch(g){}f=!1}f||(e=null)}}else"mouseover"==c?e=a.fromElement:"mouseout"==c&&(e=a.toElement);this.relatedTarget=e;d?(this.clientX=void 0!==d.clientX?d.clientX:d.pageX,this.clientY=void 0!==d.clientY?d.clientY:d.pageY,this.screenX=d.screenX||0,this.screenY=d.screenY||
0):(this.clientX=void 0!==a.clientX?a.clientX:a.pageX,this.clientY=void 0!==a.clientY?a.clientY:a.pageY,this.screenX=a.screenX||0,this.screenY=a.screenY||0);this.button=a.button;this.keyCode=a.keyCode||0;this.key=a.key||"";this.charCode=a.charCode||("keypress"==c?a.keyCode:0);this.ctrlKey=a.ctrlKey;this.altKey=a.altKey;this.shiftKey=a.shiftKey;this.metaKey=a.metaKey;this.pointerId=a.pointerId||0;this.pointerType="string"===typeof a.pointerType?a.pointerType:ce[a.pointerType]||"";this.state=a.state;
this.j=a;a.defaultPrevented&&be.A.preventDefault.call(this)};
be.prototype.stopPropagation=function(){be.A.stopPropagation.call(this);this.j.stopPropagation?this.j.stopPropagation():this.j.cancelBubble=!0};
be.prototype.preventDefault=function(){be.A.preventDefault.call(this);var a=this.j;if(a.preventDefault)a.preventDefault();else if(a.returnValue=!1,Zd)try{if(a.ctrlKey||112<=a.keyCode&&123>=a.keyCode)a.keyCode=-1}catch(b){}};var de="closure_listenable_"+(1E6*Math.random()|0);var ee=0;function fe(a,b,c,d,e){this.listener=a;this.h=null;this.src=b;this.type=c;this.capture=!!d;this.N=e;this.key=++ee;this.I=this.L=!1}
function ge(a){a.I=!0;a.listener=null;a.h=null;a.src=null;a.N=null}
;function he(a){this.src=a;this.listeners={};this.h=0}
he.prototype.add=function(a,b,c,d,e){var f=a.toString();a=this.listeners[f];a||(a=this.listeners[f]=[],this.h++);var g=ie(a,b,d,e);-1<g?(b=a[g],c||(b.L=!1)):(b=new fe(b,this.src,f,!!d,e),b.L=c,a.push(b));return b};
he.prototype.remove=function(a,b,c,d){a=a.toString();if(!(a in this.listeners))return!1;var e=this.listeners[a];b=ie(e,b,c,d);return-1<b?(ge(e[b]),Array.prototype.splice.call(e,b,1),0==e.length&&(delete this.listeners[a],this.h--),!0):!1};
function je(a,b){var c=b.type;c in a.listeners&&$a(a.listeners[c],b)&&(ge(b),0==a.listeners[c].length&&(delete a.listeners[c],a.h--))}
function ie(a,b,c,d){for(var e=0;e<a.length;++e){var f=a[e];if(!f.I&&f.listener==b&&f.capture==!!c&&f.N==d)return e}return-1}
;var ke="closure_lm_"+(1E6*Math.random()|0),le={},me=0;function ne(a,b,c,d,e){if(d&&d.once)oe(a,b,c,d,e);else if(Array.isArray(b))for(var f=0;f<b.length;f++)ne(a,b[f],c,d,e);else c=pe(c),a&&a[de]?qe(a,b,c,D(d)?!!d.capture:!!d,e):re(a,b,c,!1,d,e)}
function re(a,b,c,d,e,f){if(!b)throw Error("Invalid event type");var g=D(e)?!!e.capture:!!e,h=se(a);h||(a[ke]=h=new he(a));c=h.add(b,c,d,g,f);if(!c.h){d=te();c.h=d;d.src=a;d.listener=c;if(a.addEventListener)$d||(e=g),void 0===e&&(e=!1),a.addEventListener(b.toString(),d,e);else if(a.attachEvent)a.attachEvent(ue(b.toString()),d);else if(a.addListener&&a.removeListener)a.addListener(d);else throw Error("addEventListener and attachEvent are unavailable.");me++}}
function te(){var a=ve,b=xd?function(c){return a.call(b.src,b.listener,c)}:function(c){c=a.call(b.src,b.listener,c);
if(!c)return c};
return b}
function oe(a,b,c,d,e){if(Array.isArray(b))for(var f=0;f<b.length;f++)oe(a,b[f],c,d,e);else c=pe(c),a&&a[de]?a.h.add(String(b),c,!0,D(d)?!!d.capture:!!d,e):re(a,b,c,!0,d,e)}
function we(a,b,c,d,e){if(Array.isArray(b))for(var f=0;f<b.length;f++)we(a,b[f],c,d,e);else(d=D(d)?!!d.capture:!!d,c=pe(c),a&&a[de])?a.h.remove(String(b),c,d,e):a&&(a=se(a))&&(b=a.listeners[b.toString()],a=-1,b&&(a=ie(b,c,d,e)),(c=-1<a?b[a]:null)&&xe(c))}
function xe(a){if("number"!==typeof a&&a&&!a.I){var b=a.src;if(b&&b[de])je(b.h,a);else{var c=a.type,d=a.h;b.removeEventListener?b.removeEventListener(c,d,a.capture):b.detachEvent?b.detachEvent(ue(c),d):b.addListener&&b.removeListener&&b.removeListener(d);me--;(c=se(b))?(je(c,a),0==c.h&&(c.src=null,b[ke]=null)):ge(a)}}}
function ue(a){return a in le?le[a]:le[a]="on"+a}
function ye(a,b,c,d){var e=!0;if(a=se(a))if(b=a.listeners[b.toString()])for(b=b.concat(),a=0;a<b.length;a++){var f=b[a];f&&f.capture==c&&!f.I&&(f=ze(f,d),e=e&&!1!==f)}return e}
function ze(a,b){var c=a.listener,d=a.N||a.src;a.L&&xe(a);return c.call(d,b)}
function ve(a,b){if(a.I)return!0;if(!xd){var c=b||C("window.event"),d=new be(c,this),e=!0;if(!(0>c.keyCode||void 0!=c.returnValue)){a:{var f=!1;if(0==c.keyCode)try{c.keyCode=-1;break a}catch(k){f=!0}if(f||void 0==c.returnValue)c.returnValue=!0}c=[];for(f=d.h;f;f=f.parentNode)c.push(f);f=a.type;for(var g=c.length-1;!d.i&&0<=g;g--){d.h=c[g];var h=ye(c[g],f,!0,d);e=e&&h}for(g=0;!d.i&&g<c.length;g++)d.h=c[g],h=ye(c[g],f,!1,d),e=e&&h}return e}return ze(a,new be(b,this))}
function se(a){a=a[ke];return a instanceof he?a:null}
var Ae="__closure_events_fn_"+(1E9*Math.random()>>>0);function pe(a){if("function"===typeof a)return a;a[Ae]||(a[Ae]=function(b){return a.handleEvent(b)});
return a[Ae]}
;function K(){wc.call(this);this.h=new he(this);this.D=this;this.o=null}
F(K,wc);K.prototype[de]=!0;K.prototype.addEventListener=function(a,b,c,d){ne(this,a,b,c,d)};
K.prototype.removeEventListener=function(a,b,c,d){we(this,a,b,c,d)};
function Be(a,b){var c=a.o;if(c){var d=[];for(var e=1;c;c=c.o)d.push(c),++e}c=a.D;e=b;var f=e.type||e;if("string"===typeof e)e=new ae(e,c);else if(e instanceof ae)e.target=e.target||c;else{var g=e;e=new ae(f,c);kb(e,g)}g=!0;if(d)for(var h=d.length-1;!e.i&&0<=h;h--){var k=e.h=d[h];g=Ce(k,f,!0,e)&&g}e.i||(k=e.h=c,g=Ce(k,f,!0,e)&&g,e.i||(g=Ce(k,f,!1,e)&&g));if(d)for(h=0;!e.i&&h<d.length;h++)k=e.h=d[h],g=Ce(k,f,!1,e)&&g}
K.prototype.J=function(){K.A.J.call(this);if(this.h){var a=this.h,b=0,c;for(c in a.listeners){for(var d=a.listeners[c],e=0;e<d.length;e++)++b,ge(d[e]);delete a.listeners[c];a.h--}}this.o=null};
function qe(a,b,c,d,e){a.h.add(String(b),c,!1,d,e)}
function Ce(a,b,c,d){b=a.h.listeners[String(b)];if(!b)return!0;b=b.concat();for(var e=!0,f=0;f<b.length;++f){var g=b[f];if(g&&!g.I&&g.capture==c){var h=g.listener,k=g.N||g.src;g.L&&je(a.h,g);e=!1!==h.call(k,d)&&e}}return e&&!d.defaultPrevented}
;var De=B.JSON.stringify;function Ee(a){this.h=0;this.s=void 0;this.l=this.i=this.j=null;this.m=this.o=!1;if(a!=Ha)try{var b=this;a.call(void 0,function(c){Fe(b,2,c)},function(c){Fe(b,3,c)})}catch(c){Fe(this,3,c)}}
function Ge(){this.next=this.context=this.onRejected=this.i=this.h=null;this.j=!1}
Ge.prototype.reset=function(){this.context=this.onRejected=this.i=this.h=null;this.j=!1};
var He=new Yc(function(){return new Ge},function(a){a.reset()});
function Ie(a,b,c){var d=He.get();d.i=a;d.onRejected=b;d.context=c;return d}
Ee.prototype.then=function(a,b,c){return Je(this,"function"===typeof a?a:null,"function"===typeof b?b:null,c)};
Ee.prototype.$goog_Thenable=!0;Ee.prototype.cancel=function(a){if(0==this.h){var b=new Ke(a);fd(function(){Le(this,b)},this)}};
function Le(a,b){if(0==a.h)if(a.j){var c=a.j;if(c.i){for(var d=0,e=null,f=null,g=c.i;g&&(g.j||(d++,g.h==a&&(e=g),!(e&&1<d)));g=g.next)e||(f=g);e&&(0==c.h&&1==d?Le(c,b):(f?(d=f,d.next==c.l&&(c.l=d),d.next=d.next.next):Me(c),Ne(c,e,3,b)))}a.j=null}else Fe(a,3,b)}
function Oe(a,b){a.i||2!=a.h&&3!=a.h||Pe(a);a.l?a.l.next=b:a.i=b;a.l=b}
function Je(a,b,c,d){var e=Ie(null,null,null);e.h=new Ee(function(f,g){e.i=b?function(h){try{var k=b.call(d,h);f(k)}catch(l){g(l)}}:f;
e.onRejected=c?function(h){try{var k=c.call(d,h);void 0===k&&h instanceof Ke?g(h):f(k)}catch(l){g(l)}}:g});
e.h.j=a;Oe(a,e);return e.h}
Ee.prototype.D=function(a){this.h=0;Fe(this,2,a)};
Ee.prototype.G=function(a){this.h=0;Fe(this,3,a)};
function Fe(a,b,c){if(0==a.h){a===c&&(b=3,c=new TypeError("Promise cannot resolve to itself"));a.h=1;a:{var d=c,e=a.D,f=a.G;if(d instanceof Ee){Oe(d,Ie(e||Ha,f||null,a));var g=!0}else{if(d)try{var h=!!d.$goog_Thenable}catch(l){h=!1}else h=!1;if(h)d.then(e,f,a),g=!0;else{if(D(d))try{var k=d.then;if("function"===typeof k){Qe(d,k,e,f,a);g=!0;break a}}catch(l){f.call(a,l);g=!0;break a}g=!1}}}g||(a.s=c,a.h=b,a.j=null,Pe(a),3!=b||c instanceof Ke||Re(a,c))}}
function Qe(a,b,c,d,e){function f(k){h||(h=!0,d.call(e,k))}
function g(k){h||(h=!0,c.call(e,k))}
var h=!1;try{b.call(a,g,f)}catch(k){f(k)}}
function Pe(a){a.o||(a.o=!0,fd(a.u,a))}
function Me(a){var b=null;a.i&&(b=a.i,a.i=b.next,b.next=null);a.i||(a.l=null);return b}
Ee.prototype.u=function(){for(var a;a=Me(this);)Ne(this,a,this.h,this.s);this.o=!1};
function Ne(a,b,c,d){if(3==c&&b.onRejected&&!b.j)for(;a&&a.m;a=a.j)a.m=!1;if(b.h)b.h.j=null,Se(b,c,d);else try{b.j?b.i.call(b.context):Se(b,c,d)}catch(e){Te.call(null,e)}Zc(He,b)}
function Se(a,b,c){2==b?a.i.call(a.context,c):a.onRejected&&a.onRejected.call(a.context,c)}
function Re(a,b){a.m=!0;fd(function(){a.m&&Te.call(null,b)})}
var Te=bd;function Ke(a){Sa.call(this,a)}
F(Ke,Sa);Ke.prototype.name="cancel";function L(a){wc.call(this);this.s=1;this.l=[];this.o=0;this.h=[];this.i={};this.u=!!a}
F(L,wc);r=L.prototype;r.subscribe=function(a,b,c){var d=this.i[a];d||(d=this.i[a]=[]);var e=this.s;this.h[e]=a;this.h[e+1]=b;this.h[e+2]=c;this.s=e+3;d.push(e);return e};
function Ue(a,b,c){var d=Ve;if(a=d.i[a]){var e=d.h;(a=Za(a,function(f){return e[f+1]==b&&e[f+2]==c}))&&d.K(a)}}
r.K=function(a){var b=this.h[a];if(b){var c=this.i[b];0!=this.o?(this.l.push(a),this.h[a+1]=Ha):(c&&$a(c,a),delete this.h[a],delete this.h[a+1],delete this.h[a+2])}return!!b};
r.H=function(a,b){var c=this.i[a];if(c){for(var d=Array(arguments.length-1),e=1,f=arguments.length;e<f;e++)d[e-1]=arguments[e];if(this.u)for(e=0;e<c.length;e++){var g=c[e];We(this.h[g+1],this.h[g+2],d)}else{this.o++;try{for(e=0,f=c.length;e<f;e++)g=c[e],this.h[g+1].apply(this.h[g+2],d)}finally{if(this.o--,0<this.l.length&&0==this.o)for(;c=this.l.pop();)this.K(c)}}return 0!=e}return!1};
function We(a,b,c){fd(function(){a.apply(b,c)})}
r.clear=function(a){if(a){var b=this.i[a];b&&(H(b,this.K,this),delete this.i[a])}else this.h.length=0,this.i={}};
r.J=function(){L.A.J.call(this);this.clear();this.l.length=0};function Xe(a){this.h=a}
Xe.prototype.set=function(a,b){void 0===b?this.h.remove(a):this.h.set(a,De(b))};
Xe.prototype.get=function(a){try{var b=this.h.get(a)}catch(c){return}if(null!==b)try{return JSON.parse(b)}catch(c){throw"Storage: Invalid value was encountered";}};
Xe.prototype.remove=function(a){this.h.remove(a)};function Ye(a){this.h=a}
F(Ye,Xe);function Ze(a){this.data=a}
function $e(a){return void 0===a||a instanceof Ze?a:new Ze(a)}
Ye.prototype.set=function(a,b){Ye.A.set.call(this,a,$e(b))};
Ye.prototype.i=function(a){a=Ye.A.get.call(this,a);if(void 0===a||a instanceof Object)return a;throw"Storage: Invalid value was encountered";};
Ye.prototype.get=function(a){if(a=this.i(a)){if(a=a.data,void 0===a)throw"Storage: Invalid value was encountered";}else a=void 0;return a};function af(a){this.h=a}
F(af,Ye);af.prototype.set=function(a,b,c){if(b=$e(b)){if(c){if(c<Date.now()){af.prototype.remove.call(this,a);return}b.expiration=c}b.creation=Date.now()}af.A.set.call(this,a,b)};
af.prototype.i=function(a){var b=af.A.i.call(this,a);if(b){var c=b.creation,d=b.expiration;if(d&&d<Date.now()||c&&c>Date.now())af.prototype.remove.call(this,a);else return b}};function bf(){}
;function cf(){}
F(cf,bf);cf.prototype.clear=function(){var a=sd(this.B(!0)),b=this;H(a,function(c){b.remove(c)})};function df(a){this.h=a}
F(df,cf);r=df.prototype;r.isAvailable=function(){if(!this.h)return!1;try{return this.h.setItem("__sak","1"),this.h.removeItem("__sak"),!0}catch(a){return!1}};
r.set=function(a,b){try{this.h.setItem(a,b)}catch(c){if(0==this.h.length)throw"Storage mechanism: Storage disabled";throw"Storage mechanism: Quota exceeded";}};
r.get=function(a){a=this.h.getItem(a);if("string"!==typeof a&&null!==a)throw"Storage mechanism: Invalid value was encountered";return a};
r.remove=function(a){this.h.removeItem(a)};
r.B=function(a){var b=0,c=this.h,d=new pd;d.next=function(){if(b>=c.length)throw od;var e=c.key(b++);if(a)return e;e=c.getItem(e);if("string"!==typeof e)throw"Storage mechanism: Invalid value was encountered";return e};
return d};
r.clear=function(){this.h.clear()};
r.key=function(a){return this.h.key(a)};function ef(){var a=null;try{a=window.localStorage||null}catch(b){}this.h=a}
F(ef,df);function ff(a,b){this.i=a;this.h=null;if(Jb&&!(9<=Number(Wb))){gf||(gf=new td);this.h=gf.get(a);this.h||(b?this.h=document.getElementById(b):(this.h=document.createElement("userdata"),this.h.addBehavior("#default#userData"),document.body.appendChild(this.h)),gf.set(a,this.h));try{this.h.load(this.i)}catch(c){this.h=null}}}
F(ff,cf);var hf={".":".2E","!":".21","~":".7E","*":".2A","'":".27","(":".28",")":".29","%":"."},gf=null;function jf(a){return"_"+encodeURIComponent(a).replace(/[.!~*'()%]/g,function(b){return hf[b]})}
r=ff.prototype;r.isAvailable=function(){return!!this.h};
r.set=function(a,b){this.h.setAttribute(jf(a),b);kf(this)};
r.get=function(a){a=this.h.getAttribute(jf(a));if("string"!==typeof a&&null!==a)throw"Storage mechanism: Invalid value was encountered";return a};
r.remove=function(a){this.h.removeAttribute(jf(a));kf(this)};
r.B=function(a){var b=0,c=this.h.XMLDocument.documentElement.attributes,d=new pd;d.next=function(){if(b>=c.length)throw od;var e=c[b++];if(a)return decodeURIComponent(e.nodeName.replace(/\./g,"%")).substr(1);e=e.nodeValue;if("string"!==typeof e)throw"Storage mechanism: Invalid value was encountered";return e};
return d};
r.clear=function(){for(var a=this.h.XMLDocument.documentElement,b=a.attributes.length;0<b;b--)a.removeAttribute(a.attributes[b-1].nodeName);kf(this)};
function kf(a){try{a.h.save(a.i)}catch(b){throw"Storage mechanism: Quota exceeded";}}
;function lf(a,b){this.i=a;this.h=b+"::"}
F(lf,cf);lf.prototype.set=function(a,b){this.i.set(this.h+a,b)};
lf.prototype.get=function(a){return this.i.get(this.h+a)};
lf.prototype.remove=function(a){this.i.remove(this.h+a)};
lf.prototype.B=function(a){var b=this.i.B(!0),c=this,d=new pd;d.next=function(){for(var e=b.next();e.substr(0,c.h.length)!=c.h;)e=b.next();return a?e.substr(c.h.length):c.i.get(e)};
return d};var mf=window.yt&&window.yt.config_||window.ytcfg&&window.ytcfg.data_||{};E("yt.config_",mf);function nf(a){var b=arguments;1<b.length?mf[b[0]]=b[1]:1===b.length&&Object.assign(mf,b[0])}
function M(a,b){return a in mf?mf[a]:b}
;var of=[];function pf(a){of.forEach(function(b){return b(a)})}
function qf(a){return a&&window.yterr?function(){try{return a.apply(this,arguments)}catch(b){rf(b),pf(b)}}:a}
function rf(a){var b=C("yt.logging.errors.log");b?b(a,"ERROR",void 0,void 0,void 0):(b=M("ERRORS",[]),b.push([a,"ERROR",void 0,void 0,void 0]),nf("ERRORS",b))}
function sf(a){var b=C("yt.logging.errors.log");b?b(a,"WARNING",void 0,void 0,void 0):(b=M("ERRORS",[]),b.push([a,"WARNING",void 0,void 0,void 0]),nf("ERRORS",b))}
;var tf=0;E("ytDomDomGetNextId",C("ytDomDomGetNextId")||function(){return++tf});var uf={stopImmediatePropagation:1,stopPropagation:1,preventMouseEvent:1,preventManipulation:1,preventDefault:1,layerX:1,layerY:1,screenX:1,screenY:1,scale:1,rotation:1,webkitMovementX:1,webkitMovementY:1};
function vf(a){this.type="";this.state=this.source=this.data=this.currentTarget=this.relatedTarget=this.target=null;this.charCode=this.keyCode=0;this.metaKey=this.shiftKey=this.ctrlKey=this.altKey=!1;this.clientY=this.clientX=0;this.changedTouches=this.touches=null;try{if(a=a||window.event){this.event=a;for(var b in a)b in uf||(this[b]=a[b]);var c=a.target||a.srcElement;c&&3==c.nodeType&&(c=c.parentNode);this.target=c;var d=a.relatedTarget;if(d)try{d=d.nodeName?d:null}catch(e){d=null}else"mouseover"==
this.type?d=a.fromElement:"mouseout"==this.type&&(d=a.toElement);this.relatedTarget=d;this.clientX=void 0!=a.clientX?a.clientX:a.pageX;this.clientY=void 0!=a.clientY?a.clientY:a.pageY;this.keyCode=a.keyCode?a.keyCode:a.which;this.charCode=a.charCode||("keypress"==this.type?this.keyCode:0);this.altKey=a.altKey;this.ctrlKey=a.ctrlKey;this.shiftKey=a.shiftKey;this.metaKey=a.metaKey}}catch(e){}}
vf.prototype.preventDefault=function(){this.event&&(this.event.returnValue=!1,this.event.preventDefault&&this.event.preventDefault())};
vf.prototype.stopPropagation=function(){this.event&&(this.event.cancelBubble=!0,this.event.stopPropagation&&this.event.stopPropagation())};
vf.prototype.stopImmediatePropagation=function(){this.event&&(this.event.cancelBubble=!0,this.event.stopImmediatePropagation&&this.event.stopImmediatePropagation())};var gb=B.ytEventsEventsListeners||{};E("ytEventsEventsListeners",gb);var wf=B.ytEventsEventsCounter||{count:0};E("ytEventsEventsCounter",wf);
function xf(a,b,c,d){d=void 0===d?{}:d;a.addEventListener&&("mouseenter"!=b||"onmouseenter"in document?"mouseleave"!=b||"onmouseenter"in document?"mousewheel"==b&&"MozBoxSizing"in document.documentElement.style&&(b="MozMousePixelScroll"):b="mouseout":b="mouseover");return fb(function(e){var f="boolean"===typeof e[4]&&e[4]==!!d,g=D(e[4])&&D(d)&&hb(e[4],d);return!!e.length&&e[0]==a&&e[1]==b&&e[2]==c&&(f||g)})}
function yf(a){a&&("string"==typeof a&&(a=[a]),H(a,function(b){if(b in gb){var c=gb[b],d=c[0],e=c[1],f=c[3];c=c[4];d.removeEventListener?zf()||"boolean"===typeof c?d.removeEventListener(e,f,c):d.removeEventListener(e,f,!!c.capture):d.detachEvent&&d.detachEvent("on"+e,f);delete gb[b]}}))}
var zf=db(function(){var a=!1;try{var b=Object.defineProperty({},"capture",{get:function(){a=!0}});
window.addEventListener("test",null,b)}catch(c){}return a});
function Af(a,b,c){var d=void 0===d?{}:d;if(a&&(a.addEventListener||a.attachEvent)){var e=xf(a,b,c,d);if(!e){e=++wf.count+"";var f=!("mouseenter"!=b&&"mouseleave"!=b||!a.addEventListener||"onmouseenter"in document);var g=f?function(h){h=new vf(h);if(!jc(h.relatedTarget,function(k){return k==a}))return h.currentTarget=a,h.type=b,c.call(a,h)}:function(h){h=new vf(h);
h.currentTarget=a;return c.call(a,h)};
g=qf(g);a.addEventListener?("mouseenter"==b&&f?b="mouseover":"mouseleave"==b&&f?b="mouseout":"mousewheel"==b&&"MozBoxSizing"in document.documentElement.style&&(b="MozMousePixelScroll"),zf()||"boolean"===typeof d?a.addEventListener(b,g,d):a.addEventListener(b,g,!!d.capture)):a.attachEvent("on"+b,g);gb[e]=[a,b,c,g,d]}}}
;function Bf(a,b){"function"===typeof a&&(a=qf(a));return window.setTimeout(a,b)}
function Cf(a){"function"===typeof a&&(a=qf(a));return window.setInterval(a,250)}
;var Df=/^[\w.]*$/,Ef={q:!0,search_query:!0};function Ff(a,b){for(var c=a.split(b),d={},e=0,f=c.length;e<f;e++){var g=c[e].split("=");if(1==g.length&&g[0]||2==g.length)try{var h=Gf(g[0]||""),k=Gf(g[1]||"");h in d?Array.isArray(d[h])?cb(d[h],k):d[h]=[d[h],k]:d[h]=k}catch(p){var l=p,m=g[0],n=String(Ff);l.args=[{key:m,value:g[1],query:a,method:Hf==n?"unchanged":n}];Ef.hasOwnProperty(m)||("ReferenceError"===l.name?sf(l):rf(l))}}return d}
var Hf=String(Ff);function If(a){var b=[];eb(a,function(c,d){var e=encodeURIComponent(String(d)),f;Array.isArray(c)?f=c:f=[c];H(f,function(g){""==g?b.push(e):b.push(e+"="+encodeURIComponent(String(g)))})});
return b.join("&")}
function Jf(a){"?"==a.charAt(0)&&(a=a.substr(1));return Ff(a,"&")}
function Kf(a,b,c){var d=a.split("#",2);a=d[0];d=1<d.length?"#"+d[1]:"";var e=a.split("?",2);a=e[0];e=Jf(e[1]||"");for(var f in b)!c&&null!==e&&f in e||(e[f]=b[f]);b=a;a=pc(e);a?(c=b.indexOf("#"),0>c&&(c=b.length),f=b.indexOf("?"),0>f||f>c?(f=c,e=""):e=b.substring(f+1,c),b=[b.substr(0,f),e,b.substr(c)],c=b[1],b[1]=a?c?c+"&"+a:a:c,a=b[0]+(b[1]?"?"+b[1]:"")+b[2]):a=b;return a+d}
function Gf(a){return a&&a.match(Df)?a:decodeURIComponent(a.replace(/\+/g," "))}
;var Lf={};function Mf(a){return Lf[a]||(Lf[a]=String(a).replace(/\-([a-z])/g,function(b,c){return c.toUpperCase()}))}
;var Nf={},Of=[],Ve=new L,Pf={};function Qf(){for(var a=u(Of),b=a.next();!b.done;b=a.next())b=b.value,b()}
function Rf(a,b){var c;"yt:"==a.tagName.toLowerCase().substr(0,3)?c=a.getAttribute(b):c=a?a.dataset?a.dataset[Mf(b)]:a.getAttribute("data-"+b):null;return c}
function Sf(a,b){Ve.H.apply(Ve,arguments)}
;function Tf(a){this.i=a||{};this.j=this.h=!1;a=document.getElementById("www-widgetapi-script");if(this.h=!!("https:"==document.location.protocol||a&&0==a.src.indexOf("https:"))){a=[this.i,window.YTConfig||{}];for(var b=0;b<a.length;b++)a[b].host&&(a[b].host=a[b].host.replace("http://","https://"))}}
function N(a,b){for(var c=[a.i,window.YTConfig||{}],d=0;d<c.length;d++){var e=c[d][b];if(void 0!=e)return e}return null}
function Uf(a,b,c){Vf||(Vf={},Af(window,"message",Pa(a.l,a)));Vf[c]=b}
Tf.prototype.l=function(a){if(a.origin==N(this,"host")||a.origin==N(this,"host").replace(/^http:/,"https:")){try{var b=JSON.parse(a.data)}catch(c){return}this.j=!0;this.h||0!=a.origin.indexOf("https:")||(this.h=!0);if(a=Vf[b.id])a.u=!0,a.u&&(H(a.s,a.R,a),a.s.length=0),a.Y(b)}};
var Vf=null;function P(a){a=Wf(a);return"string"===typeof a&&"false"===a?!1:!!a}
function Xf(a,b){var c=Wf(a);return void 0===c&&void 0!==b?b:Number(c||0)}
function Wf(a){var b=M("EXPERIMENTS_FORCED_FLAGS",{});return void 0!==b[a]?b[a]:M("EXPERIMENT_FLAGS",{})[a]}
;function Yf(){}
function Zf(a,b){return $f(a,1,b)}
;function ag(){Yf.apply(this,arguments)}
x(ag,Yf);function $f(a,b,c){void 0!==c&&Number.isNaN(Number(c))&&(c=void 0);var d=C("yt.scheduler.instance.addJob");return d?d(a,b,c):void 0===c?(a(),NaN):Bf(a,c||0)}
function bg(a){if(void 0===a||!Number.isNaN(Number(a))){var b=C("yt.scheduler.instance.cancelJob");b?b(a):window.clearTimeout(a)}}
ag.prototype.start=function(){var a=C("yt.scheduler.instance.start");a&&a()};
ag.h=void 0;ag.i=function(){ag.h||(ag.h=new ag)};
ag.i();var cg=B.ytPubsubPubsubInstance||new L,dg=B.ytPubsubPubsubSubscribedKeys||{},eg=B.ytPubsubPubsubTopicToKeys||{},fg=B.ytPubsubPubsubIsSynchronous||{};L.prototype.subscribe=L.prototype.subscribe;L.prototype.unsubscribeByKey=L.prototype.K;L.prototype.publish=L.prototype.H;L.prototype.clear=L.prototype.clear;E("ytPubsubPubsubInstance",cg);E("ytPubsubPubsubTopicToKeys",eg);E("ytPubsubPubsubIsSynchronous",fg);E("ytPubsubPubsubSubscribedKeys",dg);var gg=window,R=gg.ytcsi&&gg.ytcsi.now?gg.ytcsi.now:gg.performance&&gg.performance.timing&&gg.performance.now&&gg.performance.timing.navigationStart?function(){return gg.performance.timing.navigationStart+gg.performance.now()}:function(){return(new Date).getTime()};var Ng=Xf("initial_gel_batch_timeout",1E3),Og=Math.pow(2,16)-1,Pg=null,Qg=0,Rg=void 0,Sg=0,Tg=0,Ug=0,Vg=!0,Wg=B.ytLoggingTransportGELQueue_||new Map;E("ytLoggingTransportGELQueue_",Wg);var Xg=B.ytLoggingTransportTokensToCttTargetIds_||{};E("ytLoggingTransportTokensToCttTargetIds_",Xg);function Yg(a){a=void 0===a?!1:a;return new Ee(function(b){window.clearTimeout(Sg);window.clearTimeout(Tg);Tg=0;Rg&&Rg.isReady()?(Zg(b,a),Wg.clear()):($g(),b())})}
function $g(){P("web_gel_timeout_cap")&&!Tg&&(Tg=Bf(Yg,6E4));window.clearTimeout(Sg);var a=M("LOGGING_BATCH_TIMEOUT",Xf("web_gel_debounce_ms",1E4));P("shorten_initial_gel_batch_timeout")&&Vg&&(a=Ng);Sg=Bf(Yg,a)}
function Zg(a,b){var c=Rg;b=void 0===b?!1:b;for(var d=Math.round(R()),e=Wg.size,f=u(Wg),g=f.next();!g.done;g=f.next()){var h=u(g.value);g=h.next().value;var k=h.next().value;h=ib({context:ah(c.h||bh())});h.events=k;(k=Xg[g])&&ch(h,g,k);delete Xg[g];dh(h,d);eh(c,"log_event",h,{retry:!0,onSuccess:function(){e--;e||a();Qg=Math.round(R()-d)},
onError:function(){e--;e||a()},
sa:b});Vg=!1}}
function dh(a,b){a.requestTimeMs=String(b);P("unsplit_gel_payloads_in_logs")&&(a.unsplitGelPayloadsInLogs=!0);var c=M("EVENT_ID",void 0);if(c){var d=M("BATCH_CLIENT_COUNTER",void 0)||0;!d&&P("web_client_counter_random_seed")&&(d=Math.floor(Math.random()*Og/2));d++;d>Og&&(d=1);nf("BATCH_CLIENT_COUNTER",d);c={serializedEventId:c,clientCounter:String(d)};a.serializedClientEventId=c;Pg&&Qg&&P("log_gel_rtt_web")&&(a.previousBatchInfo={serializedClientEventId:Pg,roundtripMs:String(Qg)});Pg=c;Qg=0}}
function ch(a,b,c){if(c.videoId)var d="VIDEO";else if(c.playlistId)d="PLAYLIST";else return;a.credentialTransferTokenTargetId=c;a.context=a.context||{};a.context.user=a.context.user||{};a.context.user.credentialTransferTokens=[{token:b,scope:d}]}
;var fh=B.ytLoggingGelSequenceIdObj_||{};E("ytLoggingGelSequenceIdObj_",fh);function gh(a){var b=hh;a=void 0===a?C("yt.ads.biscotti.lastId_")||"":a;var c=Object,d=c.assign,e={};e.dt=Qc;e.flash="0";a:{try{var f=b.h.top.location.href}catch(G){f=2;break a}f=f?f===b.i.location.href?0:1:2}e=(e.frm=f,e);e.u_tz=-(new Date).getTimezoneOffset();var g=void 0===g?J:g;try{var h=g.history.length}catch(G){h=0}e.u_his=h;e.u_java=!!J.navigator&&"unknown"!==typeof J.navigator.javaEnabled&&!!J.navigator.javaEnabled&&J.navigator.javaEnabled();J.screen&&(e.u_h=J.screen.height,e.u_w=J.screen.width,
e.u_ah=J.screen.availHeight,e.u_aw=J.screen.availWidth,e.u_cd=J.screen.colorDepth);J.navigator&&J.navigator.plugins&&(e.u_nplug=J.navigator.plugins.length);J.navigator&&J.navigator.mimeTypes&&(e.u_nmime=J.navigator.mimeTypes.length);h=b.h;try{var k=h.screenX;var l=h.screenY}catch(G){}try{var m=h.outerWidth;var n=h.outerHeight}catch(G){}try{var p=h.innerWidth;var q=h.innerHeight}catch(G){}k=[h.screenLeft,h.screenTop,k,l,h.screen?h.screen.availWidth:void 0,h.screen?h.screen.availTop:void 0,m,n,p,q];
l=b.h.top;try{var v=(l||window).document,w="CSS1Compat"==v.compatMode?v.documentElement:v.body;var y=(new hc(w.clientWidth,w.clientHeight)).round()}catch(G){y=new hc(-12245933,-12245933)}v=y;y={};w=new Wc;B.SVGElement&&B.document.createElementNS&&w.set(0);l=vc();l["allow-top-navigation-by-user-activation"]&&w.set(1);l["allow-popups-to-escape-sandbox"]&&w.set(2);B.crypto&&B.crypto.subtle&&w.set(3);B.TextDecoder&&B.TextEncoder&&w.set(4);w=Xc(w);y.bc=w;y.bih=v.height;y.biw=v.width;y.brdim=k.join();b=
b.i;b=(y.vis={visible:1,hidden:2,prerender:3,preview:4,unloaded:5}[b.visibilityState||b.webkitVisibilityState||b.mozVisibilityState||""]||0,y.wgl=!!J.WebGLRenderingContext,y);c=d.call(c,e,b);c.ca_type="image";a&&(c.bid=a);return c}
var hh=new function(){var a=window.document;this.h=window;this.i=a};
E("yt.ads_.signals_.getAdSignalsString",function(a){return If(gh(a))});var ih="XMLHttpRequest"in B?function(){return new XMLHttpRequest}:null;
function jh(){if(!ih)return null;var a=ih();return"open"in a?a:null}
;var kh={Authorization:"AUTHORIZATION","X-Goog-Visitor-Id":"SANDBOXED_VISITOR_ID","X-Youtube-Chrome-Connected":"CHROME_CONNECTED_HEADER","X-YouTube-Client-Name":"INNERTUBE_CONTEXT_CLIENT_NAME","X-YouTube-Client-Version":"INNERTUBE_CONTEXT_CLIENT_VERSION","X-YouTube-Delegation-Context":"INNERTUBE_CONTEXT_SERIALIZED_DELEGATION_CONTEXT","X-YouTube-Device":"DEVICE","X-Youtube-Identity-Token":"ID_TOKEN","X-YouTube-Page-CL":"PAGE_CL","X-YouTube-Page-Label":"PAGE_BUILD_LABEL","X-YouTube-Variants-Checksum":"VARIANTS_CHECKSUM"},
lh="app debugcss debugjs expflag force_ad_params force_viral_ad_response_params forced_experiments innertube_snapshots innertube_goldens internalcountrycode internalipoverride absolute_experiments conditional_experiments sbb sr_bns_address client_dev_root_url".split(" "),mh=!1;
function nh(a,b){b=void 0===b?{}:b;if(!c)var c=window.location.href;var d=a.match(kc)[1]||null,e=mc(a);d&&e?(d=c,c=a.match(kc),d=d.match(kc),c=c[3]==d[3]&&c[1]==d[1]&&c[4]==d[4]):c=e?mc(c)==e&&(Number(c.match(kc)[4]||null)||null)==(Number(a.match(kc)[4]||null)||null):!0;d=P("web_ajax_ignore_global_headers_if_set");for(var f in kh)e=M(kh[f]),!e||!c&&mc(a)||d&&void 0!==b[f]||(b[f]=e);if(c||!mc(a))b["X-YouTube-Utc-Offset"]=String(-(new Date).getTimezoneOffset());(c||!mc(a))&&(f="undefined"!=typeof Intl?
(new Intl.DateTimeFormat).resolvedOptions().timeZone:null)&&(b["X-YouTube-Time-Zone"]=f);if(c||!mc(a))b["X-YouTube-Ad-Signals"]=If(gh(void 0));return b}
function oh(a){var b=window.location.search,c=mc(a),d=lc(a.match(kc)[5]||null);d=(c=c&&(c.endsWith("youtube.com")||c.endsWith("youtube-nocookie.com")))&&d&&d.startsWith("/api/");if(!c||d)return a;var e=Jf(b),f={};H(lh,function(g){e[g]&&(f[g]=e[g])});
return Kf(a,f||{},!1)}
function ph(a,b){if(window.fetch&&"XML"!=b.format){var c={method:b.method||"GET",credentials:"same-origin"};b.headers&&(c.headers=b.headers);a=qh(a,b);var d=rh(a,b);d&&(c.body=d);b.withCredentials&&(c.credentials="include");var e=!1,f;fetch(a,c).then(function(g){if(!e){e=!0;f&&window.clearTimeout(f);var h=g.ok,k=function(l){l=l||{};var m=b.context||B;h?b.onSuccess&&b.onSuccess.call(m,l,g):b.onError&&b.onError.call(m,l,g);b.onFinish&&b.onFinish.call(m,l,g)};
"JSON"==(b.format||"JSON")&&(h||400<=g.status&&500>g.status)?g.json().then(k,function(){k(null)}):k(null)}});
b.onFetchTimeout&&0<b.timeout&&(f=Bf(function(){e||(e=!0,window.clearTimeout(f),b.onFetchTimeout.call(b.context||B))},b.timeout))}else sh(a,b)}
function sh(a,b){var c=b.format||"JSON";a=qh(a,b);var d=rh(a,b),e=!1,f=th(a,function(k){if(!e){e=!0;h&&window.clearTimeout(h);a:switch(k&&"status"in k?k.status:-1){case 200:case 201:case 202:case 203:case 204:case 205:case 206:case 304:var l=!0;break a;default:l=!1}var m=null,n=400<=k.status&&500>k.status,p=500<=k.status&&600>k.status;if(l||n||p)m=uh(a,c,k,b.convertToSafeHtml);if(l)a:if(k&&204==k.status)l=!0;else{switch(c){case "XML":l=0==parseInt(m&&m.return_code,10);break a;case "RAW":l=!0;break a}l=
!!m}m=m||{};n=b.context||B;l?b.onSuccess&&b.onSuccess.call(n,k,m):b.onError&&b.onError.call(n,k,m);b.onFinish&&b.onFinish.call(n,k,m)}},b.method,d,b.headers,b.responseType,b.withCredentials);
if(b.onTimeout&&0<b.timeout){var g=b.onTimeout;var h=Bf(function(){e||(e=!0,f.abort(),window.clearTimeout(h),g.call(b.context||B,f))},b.timeout)}}
function qh(a,b){b.includeDomain&&(a=document.location.protocol+"//"+document.location.hostname+(document.location.port?":"+document.location.port:"")+a);var c=M("XSRF_FIELD_NAME",void 0),d=b.urlParams;d&&(d[c]&&delete d[c],a=Kf(a,d||{},!0));return a}
function rh(a,b){var c=M("XSRF_FIELD_NAME",void 0),d=M("XSRF_TOKEN",void 0),e=b.postBody||"",f=b.postParams,g=M("XSRF_FIELD_NAME",void 0),h;b.headers&&(h=b.headers["Content-Type"]);b.excludeXsrf||mc(a)&&!b.withCredentials&&mc(a)!=document.location.hostname||"POST"!=b.method||h&&"application/x-www-form-urlencoded"!=h||b.postParams&&b.postParams[g]||(f||(f={}),f[c]=d);f&&"string"===typeof e&&(e=Jf(e),kb(e,f),e=b.postBodyFormat&&"JSON"==b.postBodyFormat?JSON.stringify(e):pc(e));if(!(c=e)&&(c=f)){a:{for(var k in f){f=
!1;break a}f=!0}c=!f}!mh&&c&&"POST"!=b.method&&(mh=!0,rf(Error("AJAX request with postData should use POST")));return e}
function uh(a,b,c,d){var e=null;switch(b){case "JSON":try{var f=c.responseText}catch(g){throw d=Error("Error reading responseText"),d.params=a,sf(d),g;}a=c.getResponseHeader("Content-Type")||"";f&&0<=a.indexOf("json")&&(")]}'\n"===f.substring(0,5)&&(f=f.substring(5)),e=JSON.parse(f));break;case "XML":if(a=(a=c.responseXML)?vh(a):null)e={},H(a.getElementsByTagName("*"),function(g){e[g.tagName]=wh(g)})}d&&xh(e);
return e}
function xh(a){if(D(a))for(var b in a){var c;(c="html_content"==b)||(c=b.length-5,c=0<=c&&b.indexOf("_html",c)==c);if(c){c=b;var d=a[b];if(void 0===lb){var e=null;var f=B.trustedTypes;if(f&&f.createPolicy){try{e=f.createPolicy("goog#html",{createHTML:Qa,createScript:Qa,createScriptURL:Qa})}catch(g){B.console&&B.console.error(g.message)}lb=e}else lb=e}d=(e=lb)?e.createHTML(d):d;a[c]=new Eb(d)}else xh(a[b])}}
function vh(a){return a?(a=("responseXML"in a?a.responseXML:a).getElementsByTagName("root"))&&0<a.length?a[0]:null:null}
function wh(a){var b="";H(a.childNodes,function(c){b+=c.nodeValue});
return b}
function th(a,b,c,d,e,f,g){function h(){4==(k&&"readyState"in k?k.readyState:0)&&b&&qf(b)(k)}
c=void 0===c?"GET":c;d=void 0===d?"":d;var k=jh();if(!k)return null;"onloadend"in k?k.addEventListener("loadend",h,!1):k.onreadystatechange=h;P("debug_forward_web_query_parameters")&&(a=oh(a));k.open(c,a,!0);f&&(k.responseType=f);g&&(k.withCredentials=!0);c="POST"==c&&(void 0===window.FormData||!(d instanceof FormData));if(e=nh(a,e))for(var l in e)k.setRequestHeader(l,e[l]),"content-type"==l.toLowerCase()&&(c=!1);c&&k.setRequestHeader("Content-Type","application/x-www-form-urlencoded");k.send(d);
return k}
;function yh(){return"INNERTUBE_API_KEY"in mf&&"INNERTUBE_API_VERSION"in mf}
function bh(){return{innertubeApiKey:M("INNERTUBE_API_KEY",void 0),innertubeApiVersion:M("INNERTUBE_API_VERSION",void 0),ga:M("INNERTUBE_CONTEXT_CLIENT_CONFIG_INFO"),ha:M("INNERTUBE_CONTEXT_CLIENT_NAME","WEB"),innertubeContextClientVersion:M("INNERTUBE_CONTEXT_CLIENT_VERSION",void 0),ja:M("INNERTUBE_CONTEXT_HL",void 0),ia:M("INNERTUBE_CONTEXT_GL",void 0),ka:M("INNERTUBE_HOST_OVERRIDE",void 0)||"",ma:!!M("INNERTUBE_USE_THIRD_PARTY_AUTH",!1),la:!!M("INNERTUBE_OMIT_API_KEY_WHEN_AUTH_HEADER_IS_PRESENT",
!1),appInstallData:M("SERIALIZED_CLIENT_CONFIG_DATA",void 0)}}
function ah(a){var b={client:{hl:a.ja,gl:a.ia,clientName:a.ha,clientVersion:a.innertubeContextClientVersion,configInfo:a.ga}},c=window.devicePixelRatio;c&&1!=c&&(b.client.screenDensityFloat=String(c));c=M("EXPERIMENTS_TOKEN","");""!==c&&(b.client.experimentsToken=c);c=[];var d=M("EXPERIMENTS_FORCED_FLAGS",{});for(e in d)c.push({key:e,value:String(d[e])});var e=M("EXPERIMENT_FLAGS",{});for(var f in e)f.startsWith("force_")&&void 0===d[f]&&c.push({key:f,value:String(e[f])});0<c.length&&(b.request={internalExperimentFlags:c});
a.appInstallData&&P("web_log_app_install_experiments")&&(b.client.configInfo=b.client.configInfo||{},b.client.configInfo.appInstallData=a.appInstallData);M("DELEGATED_SESSION_ID")&&!P("pageid_as_header_web")&&(b.user={onBehalfOfUser:M("DELEGATED_SESSION_ID")});a=Object;f=a.assign;e=b.client;c={};d=u(Object.entries(Jf(M("DEVICE",""))));for(var g=d.next();!g.done;g=d.next()){var h=u(g.value);g=h.next().value;h=h.next().value;"cbrand"===g?c.deviceMake=h:"cmodel"===g?c.deviceModel=h:"cbr"===g?c.browserName=
h:"cbrver"===g?c.browserVersion=h:"cos"===g?c.osName=h:"cosver"===g?c.osVersion=h:"cplatform"===g&&(c.platform=h)}b.client=f.call(a,e,c);return b}
function zh(a,b,c){c=void 0===c?{}:c;var d={"X-Goog-Visitor-Id":c.visitorData||M("VISITOR_DATA","")};if(b&&b.includes("www.youtube-nocookie.com"))return d;(b=c.Ga||M("AUTHORIZATION"))||(a?b="Bearer "+C("gapi.auth.getToken")().Fa:b=Vc([]));b&&(d.Authorization=b,d["X-Goog-AuthUser"]=M("SESSION_INDEX",0),P("pageid_as_header_web")&&(d["X-Goog-PageId"]=M("DELEGATED_SESSION_ID")));return d}
;function Ah(a){a=Object.assign({},a);delete a.Authorization;var b=Vc();if(b){var c=new md;c.update(M("INNERTUBE_API_KEY",void 0));c.update(b);b=c.digest();c=3;Ja(b);void 0===c&&(c=0);if(!bc){bc={};for(var d="ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789".split(""),e=["+/=","+/","-_=","-_.","-_"],f=0;5>f;f++){var g=d.concat(e[f].split(""));ac[f]=g;for(var h=0;h<g.length;h++){var k=g[h];void 0===bc[k]&&(bc[k]=h)}}}c=ac[c];d=[];for(e=0;e<b.length;e+=3){var l=b[e],m=(f=e+1<b.length)?
b[e+1]:0;k=(g=e+2<b.length)?b[e+2]:0;h=l>>2;l=(l&3)<<4|m>>4;m=(m&15)<<2|k>>6;k&=63;g||(k=64,f||(m=64));d.push(c[h],c[l],c[m]||"",c[k]||"")}a.hash=d.join("")}return a}
;function Bh(a){var b=new ef;(b=b.isAvailable()?a?new lf(b,a):b:null)||(a=new ff(a||"UserDataSharedStore"),b=a.isAvailable()?a:null);this.h=(a=b)?new af(a):null;this.i=document.domain||window.location.hostname}
Bh.prototype.set=function(a,b,c,d){c=c||31104E3;this.remove(a);if(this.h)try{this.h.set(a,b,Date.now()+1E3*c);return}catch(f){}var e="";if(d)try{e=escape(De(b))}catch(f){return}else e=escape(b);b=this.i;Pc.set(""+a,e,{V:c,path:"/",domain:void 0===b?"youtube.com":b,secure:!1})};
Bh.prototype.get=function(a,b){var c=void 0,d=!this.h;if(!d)try{c=this.h.get(a)}catch(e){d=!0}if(d&&(c=Pc.get(""+a,void 0))&&(c=unescape(c),b))try{c=JSON.parse(c)}catch(e){this.remove(a),c=void 0}return c};
Bh.prototype.remove=function(a){this.h&&this.h.remove(a);var b=this.i;Pc.remove(""+a,"/",void 0===b?"youtube.com":b)};var Ch;function Dh(){Ch||(Ch=new Bh("yt.innertube"));return Ch}
function Eh(a,b,c,d){if(d)return null;d=Dh().get("nextId",!0)||1;var e=Dh().get("requests",!0)||{};e[d]={method:a,request:b,authState:Ah(c),requestTime:Math.round(R())};Dh().set("nextId",d+1,86400,!0);Dh().set("requests",e,86400,!0);return d}
function Fh(a){var b=Dh().get("requests",!0)||{};delete b[a];Dh().set("requests",b,86400,!0)}
function Gh(a){var b=Dh().get("requests",!0);if(b){for(var c in b){var d=b[c];if(!(6E4>Math.round(R())-d.requestTime)){var e=d.authState,f=Ah(zh(!1));hb(e,f)&&(e=d.request,"requestTimeMs"in e&&(e.requestTimeMs=Math.round(R())),eh(a,d.method,e,{}));delete b[c]}}Dh().set("requests",b,86400,!0)}}
;var Hh=C("ytPubsub2Pubsub2Instance")||new L;L.prototype.subscribe=L.prototype.subscribe;L.prototype.unsubscribeByKey=L.prototype.K;L.prototype.publish=L.prototype.H;L.prototype.clear=L.prototype.clear;E("ytPubsub2Pubsub2Instance",Hh);E("ytPubsub2Pubsub2SubscribedKeys",C("ytPubsub2Pubsub2SubscribedKeys")||{});E("ytPubsub2Pubsub2TopicToKeys",C("ytPubsub2Pubsub2TopicToKeys")||{});E("ytPubsub2Pubsub2IsAsync",C("ytPubsub2Pubsub2IsAsync")||{});E("ytPubsub2Pubsub2SkipSubKey",null);var Ih=Yb||Zb;var Jh=[],Kh=!1;function Lh(a,b){Kh||(Jh.push({type:"EVENT",eventType:a,payload:b}),10<Jh.length&&Jh.shift())}
;function T(a,b){for(var c=[],d=1;d<arguments.length;++d)c[d-1]=arguments[d];d=Error.call(this,a);this.message=d.message;"stack"in d&&(this.stack=d.stack);d=[];var e=d.concat;if(!(c instanceof Array)){c=u(c);for(var f,g=[];!(f=c.next()).done;)g.push(f.value);c=g}this.args=e.call(d,c)}
x(T,Error);var Mh={},Nh=(Mh.AUTH_INVALID="No user identifier specified.",Mh.EXPLICIT_ABORT="Transaction was explicitly aborted.",Mh.IDB_NOT_SUPPORTED="IndexedDB is not supported.",Mh.MISSING_OBJECT_STORE="Object store not created.",Mh.UNKNOWN_ABORT="Transaction was aborted for unknown reasons.",Mh.QUOTA_EXCEEDED="The current transaction exceeded its quota limitations.",Mh.QUOTA_MAYBE_EXCEEDED="The current transaction may have failed because of exceeding quota limitations.",Mh.EXECUTE_TRANSACTION_ON_CLOSED_DB=
"Can't start a transaction on a closed database",Mh);function U(a,b,c){b=void 0===b?{}:b;c=void 0===c?Nh[a]:c;T.call(this,c,Object.assign({name:"YtIdbKnownError",isSw:void 0===self.document,isIframe:self!==self.top,type:a},b));this.type=a;this.message=c;Object.setPrototypeOf(this,U.prototype);Kh||(Jh.push({type:"ERROR",payload:this}),10<Jh.length&&Jh.shift())}
x(U,T);function Oh(a){U.call(this,"MISSING_OBJECT_STORE",{Na:a},Nh.MISSING_OBJECT_STORE);Object.setPrototypeOf(this,Oh.prototype)}
x(Oh,U);var Ph=["The database connection is closing","Can't start a transaction on a closed database","A mutation operation was attempted on a database that did not allow mutations"];
function Qh(a,b,c){if(a instanceof U||a instanceof T)return a;if("QuotaExceededError"===a.name)return new U("QUOTA_EXCEEDED",{objectStoreNames:c,dbName:b});if($b&&"UnknownError"===a.name)return new U("QUOTA_MAYBE_EXCEEDED",{objectStoreNames:c,dbName:b});if("InvalidStateError"===a.name&&Ph.some(function(d){return a.message.includes(d)}))return new U("EXECUTE_TRANSACTION_ON_CLOSED_DB",{objectStoreNames:c,
dbName:b});Object.setPrototypeOf(a,T.prototype);a.args=[{name:"IdbError",Oa:a.name,dbName:b,objectStoreNames:c}];return a}
;function Rh(a){if(!a)throw Error();throw a;}
function Sh(a){return a}
function V(a){var b=this;this.i=a;this.state={status:"PENDING"};this.h=[];this.onRejected=[];this.i(function(c){if("PENDING"===b.state.status){b.state={status:"FULFILLED",value:c};c=u(b.h);for(var d=c.next();!d.done;d=c.next())d=d.value,d()}},function(c){if("PENDING"===b.state.status){b.state={status:"REJECTED",
reason:c};c=u(b.onRejected);for(var d=c.next();!d.done;d=c.next())d=d.value,d()}})}
V.all=function(a){return new V(function(b,c){var d=[],e=a.length;0===e&&b(d);for(var f={F:0};f.F<a.length;f={F:f.F},++f.F)Th(V.resolve(a[f.F]).then(function(g){return function(h){d[g.F]=h;e--;0===e&&b(d)}}(f)),function(g){c(g)})})};
V.resolve=function(a){return new V(function(b,c){a instanceof V?a.then(b,c):b(a)})};
V.reject=function(a){return new V(function(b,c){c(a)})};
V.prototype.then=function(a,b){var c=this,d=null!==a&&void 0!==a?a:Sh,e=null!==b&&void 0!==b?b:Rh;return new V(function(f,g){"PENDING"===c.state.status?(c.h.push(function(){Uh(c,c,d,f,g)}),c.onRejected.push(function(){Vh(c,c,e,f,g)})):"FULFILLED"===c.state.status?Uh(c,c,d,f,g):"REJECTED"===c.state.status&&Vh(c,c,e,f,g)})};
function Th(a,b){a.then(void 0,b)}
function Uh(a,b,c,d,e){try{if("FULFILLED"!==a.state.status)throw Error("calling handleResolve before the promise is fulfilled.");var f=c(a.state.value);f instanceof V?Wh(a,b,f,d,e):d(f)}catch(g){e(g)}}
function Vh(a,b,c,d,e){try{if("REJECTED"!==a.state.status)throw Error("calling handleReject before the promise is rejected.");var f=c(a.state.reason);f instanceof V?Wh(a,b,f,d,e):d(f)}catch(g){e(g)}}
function Wh(a,b,c,d,e){b===c?e(new TypeError("Circular promise chain detected.")):c.then(function(f){f instanceof V?Wh(a,b,f,d,e):d(f)},function(f){e(f)})}
;function Xh(a,b,c){function d(){c(a.error);f()}
function e(){b(a.result);f()}
function f(){try{a.removeEventListener("success",e),a.removeEventListener("error",d)}catch(g){}}
a.addEventListener("success",e);a.addEventListener("error",d)}
function Yh(a){return new Promise(function(b,c){Xh(a,b,c)})}
function W(a){return new V(function(b,c){Xh(a,b,c)})}
;function Zh(a,b){return new V(function(c,d){function e(){var f=a?b(a):null;f?f.then(function(g){a=g;e()},d):c()}
e()})}
;function $h(a,b){this.h=a;this.options=b;this.transactionCount=0;this.j=Math.round(R());this.i=!1}
r=$h.prototype;r.add=function(a,b,c){return ai(this,[a],"readwrite",function(d){return bi(d,a).add(b,c)})};
r.clear=function(a){return ai(this,[a],"readwrite",function(b){return bi(b,a).clear()})};
r.close=function(){var a;this.h.close();(null===(a=this.options)||void 0===a?0:a.closed)&&this.options.closed()};
r.count=function(a,b){return ai(this,[a],"readonly",function(c){return bi(c,a).count(b)})};
r["delete"]=function(a,b){return ai(this,[a],"readwrite",function(c){return bi(c,a)["delete"](b)})};
r.get=function(a,b){return ai(this,[a],"readwrite",function(c){return bi(c,a).get(b)})};
function ai(a,b,c,d){c=void 0===c?"readonly":c;a.transactionCount++;try{var e=a.h.transaction(b,c);var f=ci(e,d)["catch"](function(g){throw Qh(g,a.h.name,b.join());})}catch(g){f=g instanceof Error?Promise.reject(Qh(g,a.h.name,b.join())):Promise.reject(Qh(Error("unexpected transaction error: "+g),a.h.name,b.join()))}di(a,f,b.join(),c);
return f}
function di(a,b,c,d){Ra(a,function f(){var g,h,k=this,l,m,n;return Ba(f,function(p){if(1==p.h)return g=Math.round(R()),p.l=2,z(p,b,4);2!=p.h?(h=Math.round(R()),ei(k,!0,c,h-g),p.h=0,p.l=0):(l=va(p),m=Math.round(R()),n=m-g,l instanceof U&&("QUOTA_EXCEEDED"===l.type||"QUOTA_MAYBE_EXCEEDED"===l.type)&&Lh("QUOTA_EXCEEDED",{dbName:k.h.name,objectStoreNames:c,transactionCount:k.transactionCount,transactionMode:d}),l instanceof U&&"UNKNOWN_ABORT"===l.type&&(Lh("TRANSACTION_UNEXPECTEDLY_ABORTED",{objectStoreNames:c,
transactionDuration:n,transactionCount:k.transactionCount,dbDuration:m-k.j}),k.i=!0),ei(k,!1,c,n),p.h=0)})})}
function ei(a,b,c,d){Lh("TRANSACTION_ENDED",{objectStoreNames:c,connectionHasUnknownAbortedTransaction:a.i,duration:d,isSuccessful:b})}
function fi(a){this.h=a}
r=fi.prototype;r.add=function(a,b){return W(this.h.add(a,b))};
r.clear=function(){return W(this.h.clear()).then(function(){})};
r.count=function(a){return W(this.h.count(a))};
function gi(a,b){return hi(a,{query:b},function(c){return c["delete"]().then(function(){return c["continue"]()})}).then(function(){})}
r["delete"]=function(a){return a instanceof IDBKeyRange?gi(this,a):W(this.h["delete"](a))};
r.get=function(a){return W(this.h.get(a))};
r.index=function(a){return new ii(this.h.index(a))};
r.getName=function(){return this.h.name};
function hi(a,b,c){a=a.h.openCursor(b.query,b.direction);return ji(a).then(function(d){return Zh(d,c)})}
function ki(a){var b=this;this.h=a;this.i=new Map;this.aborted=!1;this.done=new Promise(function(c,d){b.h.addEventListener("complete",function(){c()});
b.h.addEventListener("error",function(e){e.currentTarget===e.target&&d(b.h.error)});
b.h.addEventListener("abort",function(){var e=b.h.error;if(e)d(e);else if(!b.aborted){e=U;for(var f=b.h.objectStoreNames,g=[],h=0;h<f.length;h++){var k=f.item(h);if(null===k)throw Error("Invariant: item in DOMStringList is null");g.push(k)}e=new e("UNKNOWN_ABORT",{objectStoreNames:g.join(),dbName:b.h.db.name,mode:b.h.mode});d(e)}})})}
function ci(a,b){var c=new ki(a);return li(c,b)}
function li(a,b){var c=new Promise(function(d,e){Th(b(a).then(function(f){a.commit();d(f)}),e)});
return Promise.all([c,a.done]).then(function(d){return u(d).next().value})}
ki.prototype.abort=function(){this.h.abort();this.aborted=!0;throw new U("EXPLICIT_ABORT");};
ki.prototype.commit=function(){var a=this.h;a.commit&&!this.aborted&&a.commit()};
function bi(a,b){var c=a.h.objectStore(b),d=a.i.get(c);d||(d=new fi(c),a.i.set(c,d));return d}
function ii(a){this.h=a}
ii.prototype.count=function(a){return W(this.h.count(a))};
ii.prototype["delete"]=function(a){return mi(this,{query:a},function(b){return b["delete"]().then(function(){return b["continue"]()})})};
ii.prototype.get=function(a){return W(this.h.get(a))};
ii.prototype.getKey=function(a){return W(this.h.getKey(a))};
function mi(a,b,c){a=a.h.openCursor(void 0===b.query?null:b.query,void 0===b.direction?"next":b.direction);return ji(a).then(function(d){return Zh(d,c)})}
function ni(a,b){this.request=a;this.cursor=b}
function ji(a){return W(a).then(function(b){return null===b?null:new ni(a,b)})}
r=ni.prototype;r.advance=function(a){this.cursor.advance(a);return ji(this.request)};
r["continue"]=function(a){this.cursor["continue"](a);return ji(this.request)};
r["delete"]=function(){return W(this.cursor["delete"]()).then(function(){})};
r.getKey=function(){return this.cursor.key};
r.update=function(a){return W(this.cursor.update(a))};function oi(a,b,c){return Ra(this,function e(){var f,g,h,k,l,m,n,p,q,v;return Ba(e,function(w){if(1==w.h)return f=self.indexedDB.open(a,b),g=c,h=g.blocked,k=g.blocking,l=g.ra,m=g.upgrade,n=g.closed,q=function(){p||(p=new $h(f.result,{closed:n}));return p},f.addEventListener("upgradeneeded",function(y){if(null===y.newVersion)throw Error("Invariant: newVersion on IDbVersionChangeEvent is null");
if(null===f.transaction)throw Error("Invariant: transaction on IDbOpenDbRequest is null");y.dataLoss&&"none"!==y.dataLoss&&Lh("IDB_DATA_CORRUPTED",{reason:y.dataLossMessage||"unknown reason",dbName:a});var G=q(),Q=new ki(f.transaction);m&&m(G,y.oldVersion,y.newVersion,Q)}),h&&f.addEventListener("blocked",function(){h()}),z(w,Yh(f),2);
v=w.m;k&&v.addEventListener("versionchange",function(){k(q())});
v.addEventListener("close",function(){Lh("IDB_UNEXPECTEDLY_CLOSED",{dbName:a,dbVersion:v.version});l&&l()});
return w["return"](q())})})}
function pi(a,b){b=void 0===b?{}:b;return Ra(this,function d(){var e,f,g;return Ba(d,function(h){e=self.indexedDB.deleteDatabase(a);f=b;(g=f.blocked)&&e.addEventListener("blocked",function(){g()});
return z(h,Yh(e),0)})})}
;function qi(a){this.name="YtIdbMeta";this.options=a;this.i=!1}
function ri(a,b,c){c=void 0===c?{}:c;c=void 0===c?{}:c;return oi(a,b,c)}
qi.prototype["delete"]=function(a){a=void 0===a?{}:a;return pi(this.name,a)};
function si(){var a=ti;if(!a.h){var b=function(){a.h===e&&(a.h=void 0)},c={blocking:function(f){f.close()},
closed:b,ra:b,upgrade:a.options.upgrade},d=function(){return Ra(a,function g(){var h=this,k,l,m;return Ba(g,function(n){switch(n.h){case 1:return n.l=2,z(n,ri(h.name,h.options.version,c),4);case 4:k=n.m;if(!Xb){n.h=5;break}a:{var p=u(Object.keys(h.options.qa));for(var q=p.next();!q.done;q=p.next())if(q=q.value,!k.h.objectStoreNames.contains(q)){p=q;break a}p=void 0}l=p;if(void 0===l){n.h=5;break}if(!Xb||h.i){n.h=7;break}h.i=!0;return z(n,h["delete"](),8);case 8:return n["return"](d());case 7:throw new Oh(l);
case 5:return n["return"](k);case 2:m=va(n);if(m instanceof DOMException?"VersionError"===m.name:"DOMError"in self&&m instanceof DOMError?"VersionError"===m.name:m instanceof Object&&"message"in m&&"An attempt was made to open a database using a lower version than the existing version."===m.message)return n["return"](ri(h.name,void 0,Object.assign(Object.assign({},c),{upgrade:void 0})));b();throw m;}})})};
var e=d();a.h=e}return a.h}
;var ti=new qi({qa:{databases:!0},upgrade:function(a,b){1>b&&a.h.createObjectStore("databases",{keyPath:"actualName"})}});
function ui(a){return Ra(this,function c(){var d;return Ba(c,function(e){if(1==e.h)return z(e,si(),2);d=e.m;return e["return"](ai(d,["databases"],"readwrite",function(f){var g=bi(f,"databases");return g.get(a.actualName).then(function(h){if(h?a.actualName!==h.actualName||a.publicName!==h.publicName||a.userIdentifier!==h.userIdentifier||a.signedIn!==h.signedIn||a.clearDataOnAuthChange!==h.clearDataOnAuthChange:1)return W(g.h.put(a,void 0)).then(function(){})})}))})})}
function vi(){return Ra(this,function b(){var c;return Ba(b,function(d){if(1==d.h)return z(d,si(),2);c=d.m;return d["return"](c["delete"]("databases","yt-idb-test-do-not-use"))})})}
;var wi;
function xi(){return Ra(this,function b(){var c,d;return Ba(b,function(e){switch(e.h){case 1:var f;if(f=Ih)f=/WebKit\/([0-9]+)/.exec(Bb),f=!!(f&&600<=parseInt(f[1],10));f&&(f=/WebKit\/([0-9]+)/.exec(Bb),f=!(f&&602<=parseInt(f[1],10)));if(f&&!P("ytidb_allow_on_ios_safari_v8_and_v9")||Kb)return e["return"](!1);try{if(c=self,!(c.indexedDB&&c.IDBIndex&&c.IDBKeyRange&&c.IDBObjectStore))return e["return"](!1)}catch(g){return e["return"](!1)}if(!("IDBTransaction"in self&&"objectStoreNames"in IDBTransaction.prototype))return e["return"](!1);e.l=
2;d={actualName:"yt-idb-test-do-not-use",publicName:"yt-idb-test-do-not-use",userIdentifier:void 0,signedIn:!1};return z(e,ui(d),4);case 4:return z(e,vi(),5);case 5:return e["return"](!0);case 2:return va(e),e["return"](!1)}})})}
function yi(){if(void 0!==wi)return wi;Kh=!0;return wi=xi().then(function(a){Kh=!1;return a})}
;var zi;function Ai(){zi||(zi=new Bh("yt.offline"));return zi}
;function Bi(){K.call(this);this.s=this.u=this.i=!1;this.l=Ci();Di(this);Ei(this)}
x(Bi,K);function Ci(){var a=window.navigator.onLine;return void 0===a?!0:a}
function Ei(a){window.addEventListener("online",function(){a.l=!0;a.i&&Be(a,"ytnetworkstatus-online");Fi(a);if(a.s&&P("offline_error_handling")){var b=Ai().get("errors",!0);if(b){for(var c in b)if(b[c]){var d=new T(c,"sent via offline_errors");d.name=b[c].name;d.stack=b[c].stack;rf(d)}Ai().set("errors",{},2592E3,!0)}}})}
function Di(a){window.addEventListener("offline",function(){a.l=!1;a.i&&Be(a,"ytnetworkstatus-offline");Fi(a)})}
function Fi(a){a.u&&(sf(new T("NetworkStatusManager state did not match poll",R()-0)),a.u=!1)}
;function Gi(a){a=void 0===a?{}:a;K.call(this);var b=this;this.l=this.u=0;Bi.h||(Bi.h=new Bi);this.i=Bi.h;this.i.i=!0;a.oa&&(this.i.s=!0);a.O?(this.O=a.O,qe(this.i,"ytnetworkstatus-online",function(){Hi(b,"publicytnetworkstatus-online")}),qe(this.i,"ytnetworkstatus-offline",function(){Hi(b,"publicytnetworkstatus-offline")})):(qe(this.i,"ytnetworkstatus-online",function(){Be(b,"publicytnetworkstatus-online")}),qe(this.i,"ytnetworkstatus-offline",function(){Be(b,"publicytnetworkstatus-offline")}))}
x(Gi,K);function Hi(a,b){a.O?a.l?(bg(a.u),a.u=Zf(function(){a.s!==b&&(Be(a,b),a.s=b,a.l=R())},a.O-(R()-a.l))):(Be(a,b),a.s=b,a.l=R()):Be(a,b)}
;var Ii;function Ji(a,b){b=void 0===b?{}:b;yi().then(function(){Ii||(Ii=new Gi({oa:!0}));Ii.i.l!==Ci()&&sf(new T("NetworkStatusManager isOnline does not match window status"));sh(a,b)})}
function Ki(a,b){b=void 0===b?{}:b;yi().then(function(){sh(a,b)})}
;function Li(a){var b=this;this.h=null;a?this.h=a:yh()&&(this.h=bh());$f(function(){Gh(b)},0,5E3)}
Li.prototype.isReady=function(){!this.h&&yh()&&(this.h=bh());return!!this.h};
function eh(a,b,c,d){!M("VISITOR_DATA")&&"visitor_id"!==b&&.01>Math.random()&&sf(new T("Missing VISITOR_DATA when sending innertube request.",b,c,d));if(!a.isReady()){var e=new T("innertube xhrclient not ready",b,c,d);rf(e);throw e;}var f={headers:{"Content-Type":"application/json"},method:"POST",postParams:c,postBodyFormat:"JSON",onTimeout:function(){d.onTimeout()},
onFetchTimeout:d.onTimeout,onSuccess:function(n,p){if(d.onSuccess)d.onSuccess(p)},
onFetchSuccess:function(n){if(d.onSuccess)d.onSuccess(n)},
onError:function(n,p){if(d.onError)d.onError(p)},
onFetchError:function(n){if(d.onError)d.onError(n)},
timeout:d.timeout,withCredentials:!0},g="";(e=a.h.ka)&&(g=e);var h=a.h.ma||!1,k=zh(h,g,d);Object.assign(f.headers,k);f.headers.Authorization&&!g&&(f.headers["x-origin"]=window.location.origin);e="/youtubei/"+a.h.innertubeApiVersion+"/"+b;var l={alt:"json"};a.h.la&&f.headers.Authorization||(l.key=a.h.innertubeApiKey);var m=Kf(""+g+e,l||{},!0);yi().then(function(n){if(d.retry&&P("retry_web_logging_batches")&&"www.youtube-nocookie.com"!=g){if(P("networkless_gel")&&!n||!P("networkless_gel"))var p=Eh(b,
c,k,h);if(p){var q=f.onSuccess,v=f.onFetchSuccess;f.onSuccess=function(w,y){Fh(p);q(w,y)};
c.onFetchSuccess=function(w,y){Fh(p);v(w,y)}}}try{P("use_fetch_for_op_xhr")?ph(m,f):P("networkless_gel")&&d.retry?(f.method="POST",!d.sa&&P("nwl_send_fast_on_unload")?Ki(m,f):Ji(m,f)):(f.method="POST",f.postParams||(f.postParams={}),sh(m,f))}catch(w){if("InvalidAccessError"==w.name)p&&(Fh(p),p=0),sf(Error("An extension is blocking network request."));
else throw w;}p&&$f(function(){Gh(a)},0,5E3)})}
;function Mi(a,b){var c=void 0===c?{}:c;var d=Li;M("ytLoggingEventsDefaultDisabled",!1)&&Li==Li&&(d=null);var e=c;e=void 0===e?{}:e;c={};c.eventTimeMs=Math.round(e.timestamp||R());c[a]=b;var f=C("_lact",window);f=null==f?-1:Math.max(Date.now()-f,0);c.context={lastActivityMs:String(e.timestamp||!isFinite(f)?-1:f)};if(P("log_sequence_info_on_gel_web")&&e.Z){f=c.context;var g=e.Z;fh[g]=g in fh?fh[g]+1:0;f.sequence={index:fh[g],groupKey:g};e.La&&delete fh[e.Z]}e=e.Ka;f="";e&&(f={},e.videoId?f.videoId=
e.videoId:e.playlistId&&(f.playlistId=e.playlistId),Xg[e.token]=f,f=e.token);e=Wg.get(f)||[];Wg.set(f,e);e.push(c);d&&(Rg=new d);d=Xf("web_logging_max_batch")||100;c=R();e.length>=d?Yg(!0):10<=c-Ug&&($g(),Ug=c)}
;var Ni=[{W:function(a){return"Cannot read property '"+a.key+"'"},
P:{TypeError:[{regexp:/Cannot read property '([^']+)' of (null|undefined)/,groups:["key","value"]},{regexp:/\u65e0\u6cd5\u83b7\u53d6\u672a\u5b9a\u4e49\u6216 (null|undefined) \u5f15\u7528\u7684\u5c5e\u6027\u201c([^\u201d]+)\u201d/,groups:["value","key"]},{regexp:/\uc815\uc758\ub418\uc9c0 \uc54a\uc74c \ub610\ub294 (null|undefined) \ucc38\uc870\uc778 '([^']+)' \uc18d\uc131\uc744 \uac00\uc838\uc62c \uc218 \uc5c6\uc2b5\ub2c8\ub2e4./,groups:["value","key"]},{regexp:/No se puede obtener la propiedad '([^']+)' de referencia nula o sin definir/,
groups:["key"]},{regexp:/Unable to get property '([^']+)' of (undefined or null) reference/,groups:["key","value"]}],Error:[{regexp:/(Permission denied) to access property "([^']+)"/,groups:["reason","key"]}]}},{W:function(a){return"Cannot call '"+a.key+"'"},
P:{TypeError:[{regexp:/(?:([^ ]+)?\.)?([^ ]+) is not a function/,groups:["base","key"]},{regexp:/([^ ]+) called on (null or undefined)/,groups:["key","value"]},{regexp:/Object (.*) has no method '([^ ]+)'/,groups:["base","key"]},{regexp:/Object doesn't support property or method '([^ ]+)'/,groups:["key"]},{regexp:/\u30aa\u30d6\u30b8\u30a7\u30af\u30c8\u306f '([^']+)' \u30d7\u30ed\u30d1\u30c6\u30a3\u307e\u305f\u306f\u30e1\u30bd\u30c3\u30c9\u3092\u30b5\u30dd\u30fc\u30c8\u3057\u3066\u3044\u307e\u305b\u3093/,
groups:["key"]},{regexp:/\uac1c\uccb4\uac00 '([^']+)' \uc18d\uc131\uc774\ub098 \uba54\uc11c\ub4dc\ub97c \uc9c0\uc6d0\ud558\uc9c0 \uc54a\uc2b5\ub2c8\ub2e4./,groups:["key"]}]}}];function Oi(){this.h=[];this.i=[]}
var Pi;function Qi(){Pi||(Pi=new Oi);return Pi}
;var Ri=new L;function Si(a){function b(){return a.charCodeAt(d++)}
var c=a.length,d=0;do{var e=Ti(b);if(Infinity===e)break;var f=e>>3;switch(e&7){case 0:e=Ti(b);if(2===f)return e;break;case 1:if(2===f)return;d+=8;break;case 2:e=Ti(b);if(2===f)return a.substr(d,e);d+=e;break;case 5:if(2===f)return;d+=4;break;default:return}}while(d<c)}
function Ti(a){var b=a(),c=b&127;if(128>b)return c;b=a();c|=(b&127)<<7;if(128>b)return c;b=a();c|=(b&127)<<14;if(128>b)return c;b=a();return 128>b?c|(b&127)<<21:Infinity}
;function Ui(a,b,c,d){c+="."+a;a=Vi(b);d[c]=a;return c.length+a.length}
function Vi(a){return("string"===typeof a?a:String(JSON.stringify(a))).substr(0,500)}
;var Wi=new Set,Xi=0,Yi=0,Zi=0,$i=["PhantomJS","Googlebot","TO STOP THIS SECURITY SCAN go/scan"];function Y(a,b,c){this.o=this.h=this.i=null;this.m=Ka(this);this.j=0;this.u=!1;this.s=[];this.l=null;this.D=c;this.G={};c=document;if(a="string"===typeof a?c.getElementById(a):a)if(c="iframe"==a.tagName.toLowerCase(),b.host||(b.host=c?nc(a.src):"https://www.youtube.com"),this.i=new Tf(b),c||(b=aj(this,a),this.o=a,(c=a.parentNode)&&c.replaceChild(b,a),a=b),this.h=a,this.h.id||(this.h.id="widget"+Ka(this.h)),Nf[this.h.id]=this,window.postMessage){this.l=new L;bj(this);b=N(this.i,"events");for(var d in b)b.hasOwnProperty(d)&&
this.addEventListener(d,b[d]);for(var e in Pf)cj(this,e)}}
r=Y.prototype;r.setSize=function(a,b){this.h.width=a;this.h.height=b;return this};
r.pa=function(){return this.h};
r.Y=function(a){dj(this,a.event,a)};
r.addEventListener=function(a,b){var c=b;"string"==typeof b&&(c=function(){window[b].apply(window,arguments)});
if(!c)return this;this.l.subscribe(a,c);ej(this,a);return this};
function cj(a,b){var c=b.split(".");if(2==c.length){var d=c[1];a.D==c[0]&&ej(a,d)}}
r.destroy=function(){this.h.id&&(Nf[this.h.id]=null);var a=this.l;a&&"function"==typeof a.dispose&&a.dispose();if(this.o){a=this.h;var b=a.parentNode;b&&b.replaceChild(this.o,a)}else(a=this.h)&&a.parentNode&&a.parentNode.removeChild(a);Vf&&(Vf[this.m]=null);this.i=null;a=this.h;for(var c in gb)gb[c][0]==a&&yf(c);this.o=this.h=null};
r.T=function(){return{}};
function fj(a,b,c){c=c||[];c=Array.prototype.slice.call(c);b={event:"command",func:b,args:c};a.u?a.R(b):a.s.push(b)}
function dj(a,b,c){a.l.j||(c={target:a,data:c},a.l.H(b,c),Sf(a.D+"."+b,c))}
function aj(a,b){for(var c=document.createElement("iframe"),d=b.attributes,e=0,f=d.length;e<f;e++){var g=d[e].value;null!=g&&""!=g&&"null"!=g&&c.setAttribute(d[e].name,g)}c.setAttribute("frameBorder",0);c.setAttribute("allowfullscreen",1);c.setAttribute("allow","accelerometer; autoplay; clipboard-write; encrypted-media; gyroscope; picture-in-picture");c.setAttribute("title","YouTube "+N(a.i,"title"));(d=N(a.i,"width"))&&c.setAttribute("width",d);(d=N(a.i,"height"))&&c.setAttribute("height",d);var h=
a.T();h.enablejsapi=window.postMessage?1:0;window.location.host&&(h.origin=window.location.protocol+"//"+window.location.host);h.widgetid=a.m;window.location.href&&H(["debugjs","debugcss"],function(k){var l=window.location.href;var m=l.search(qc);b:{var n=0;for(var p=k.length;0<=(n=l.indexOf(k,n))&&n<m;){var q=l.charCodeAt(n-1);if(38==q||63==q)if(q=l.charCodeAt(n+p),!q||61==q||38==q||35==q)break b;n+=p+1}n=-1}if(0>n)l=null;else{p=l.indexOf("&",n);if(0>p||p>m)p=m;n+=k.length+1;l=decodeURIComponent(l.substr(n,
p-n).replace(/\+/g," "))}null!==l&&(h[k]=l)});
c.src=N(a.i,"host")+("/embed/"+N(a.i,"videoId"))+"?"+pc(h);return c}
r.X=function(){this.h&&this.h.contentWindow?this.R({event:"listening"}):window.clearInterval(this.j)};
function bj(a){Uf(a.i,a,a.m);a.j=Cf(Pa(a.X,a));Af(a.h,"load",Pa(function(){window.clearInterval(this.j);this.j=Cf(Pa(this.X,this))},a))}
function ej(a,b){a.G[b]||(a.G[b]=!0,fj(a,"addEventListener",[b]))}
r.R=function(a){a.id=this.m;a.channel="widget";var b=De(a),c=this.i,d=nc(this.h.src||"");var e=0==d.indexOf("https:")?[d]:c.h?[d.replace("http:","https:")]:c.j?[d]:[d,d.replace("http:","https:")];if(this.h.contentWindow)for(var f=0;f<e.length;f++)try{this.h.contentWindow.postMessage(b,e[f])}catch(Ua){if(Ua.name&&"SyntaxError"==Ua.name){if(!(Ua.message&&0<Ua.message.indexOf("target origin ''"))){var g=void 0,h=Ua;g=void 0===g?{}:g;g.name=M("INNERTUBE_CONTEXT_CLIENT_NAME",1);g.version=M("INNERTUBE_CONTEXT_CLIENT_VERSION",
void 0);var k=g||{},l="WARNING";l=void 0===l?"ERROR":l;if(h){if(P("console_log_js_exceptions")){var m=h,n=[];n.push("Name: "+m.name);n.push("Message: "+m.message);m.hasOwnProperty("params")&&n.push("Error Params: "+JSON.stringify(m.params));n.push("File name: "+m.fileName);n.push("Stacktrace: "+m.stack);window.console.log(n.join("\n"),m)}if(!(5<=Xi)){var p=h,q=k,v=Kc(p),w=v.message||"Unknown Error",y=v.name||"UnknownError",G=v.stack||p.h||"Not available";if(G.startsWith(y+": "+w)){var Q=G.split("\n");
Q.shift();G=Q.join("\n")}var Ia=v.lineNumber||"Not available",kj=v.fileName||"Not available",lj=G,cc=void 0,ha=p,S=q;if(ha.hasOwnProperty("args")&&ha.args&&ha.args.length)for(var sa=0,dc=0;dc<ha.args.length;dc++){var O=ha.args[dc],ta="params."+dc;sa+=ta.length;if(O)if(Array.isArray(O)){for(var mj=S,yd=sa,sb=0;sb<O.length&&!(O[sb]&&(yd+=Ui(sb,O[sb],ta,mj),500<yd));sb++);sa=yd}else if("object"===typeof O){var Va=void 0,hg=S,tb=sa;for(Va in O)if(O[Va]){var nj=tb;var zd=Va,ig=O[Va],oj=hg;if("string"!==
typeof ig||"clickTrackingParams"!==zd&&"trackingParams"!==zd)var jg=0;else{var kg=Si(atob(ig.replace(/-/g,"+").replace(/_/g,"/")));jg=kg?Ui(zd+".ve",kg,ta,oj):0}tb=nj+jg;tb+=Ui(Va,O[Va],ta,hg);if(500<tb)break}sa=tb}else S[ta]=Vi(O),sa+=S[ta].length;else S[ta]=Vi(O),sa+=S[ta].length;if(500<=sa)break}else if(ha.hasOwnProperty("params")&&ha.params){var ec=ha.params;if("object"===typeof ha.params){var lg=0;for(cc in ec)if(ec[cc]){var mg="params."+cc,ng=Vi(ec[cc]);S[mg]=ng;lg+=mg.length+ng.length;if(500<
lg)break}}else S.params=Vi(ec)}navigator.vendor&&!S.hasOwnProperty("vendor")&&(S["device.vendor"]=navigator.vendor);var ub={message:w,name:y,lineNumber:Ia,fileName:kj,stack:lj,params:S,sampleWeight:1},og=Number(p.columnNumber);isNaN(og)||(ub.lineNumber=ub.lineNumber+":"+og);if(void 0!==p.sampleWeight)var fc=p.sampleWeight;else a:{for(var Ad=ub,pg=Qi(),qg=u(pg.i),Bd=qg.next();!Bd.done;Bd=qg.next()){var rg=Bd.value;if(Ad.message&&Ad.message.match(rg.Ma)){fc=rg.weight;break a}}for(var sg=u(pg.h),Cd=
sg.next();!Cd.done;Cd=sg.next()){var tg=Cd.value;if(tg.Ia(Ad)){fc=tg.weight;break a}}fc=1}ub.sampleWeight=fc;for(var Ea=ub,ug=u(Ni),Dd=ug.next();!Dd.done;Dd=ug.next()){var Ed=Dd.value;if(Ed.P[Ea.name])for(var vg=u(Ed.P[Ea.name]),Fd=vg.next();!Fd.done;Fd=vg.next()){var wg=Fd.value,gc=Ea.message.match(wg.regexp);if(gc){Ea.params["params.error.original"]=gc[0];for(var Gd=wg.groups,xg={},Wa=0;Wa<Gd.length;Wa++)xg[Gd[Wa]]=gc[Wa+1],Ea.params["params.error."+Gd[Wa]]=gc[Wa+1];Ea.message=Ed.W(xg);break}}}var Xa=
Ea;Xa.params||(Xa.params={});var yg=Qi();Xa.params["params.errorServiceSignature"]="msg="+yg.i.length+"&cb="+yg.h.length;Xa.params["params.serviceWorker"]="false";Xa.params["params.fscripts"]=String(document.querySelectorAll("script:not([nonce])").length);var ia=Xa;window.yterr&&"function"===typeof window.yterr&&window.yterr(ia);var pj=0===ia.sampleWeight;if(!Wi.has(ia.message)&&!pj){"ERROR"===l?(Ri.H("handleError",ia),P("record_app_crashed_web")&&0===Zi&&1===ia.sampleWeight&&(Zi++,Mi("appCrashed",
{appCrashType:"APP_CRASH_TYPE_BREAKPAD"})),Yi++):"WARNING"===l&&Ri.H("handleWarning",ia);if(P("kevlar_gel_error_routing")){var zg=l,X=ia;a:{for(var Ag=u($i),Hd=Ag.next();!Hd.done;Hd=Ag.next()){var Bg=Bb;if(Bg&&0<=Bg.toLowerCase().indexOf(Hd.value.toLowerCase())){var Cg=!0;break a}}Cg=!1}if(!Cg){var vb={stackTrace:X.stack};X.fileName&&(vb.filename=X.fileName);var ja=X.lineNumber&&X.lineNumber.split?X.lineNumber.split(":"):[];0!==ja.length&&(1!==ja.length||isNaN(Number(ja[0]))?2!==ja.length||isNaN(Number(ja[0]))||
isNaN(Number(ja[1]))||(vb.lineNumber=Number(ja[0]),vb.columnNumber=Number(ja[1])):vb.lineNumber=Number(ja[0]));var Id={level:"ERROR_LEVEL_UNKNOWN",message:X.message,errorClassName:X.name,sampleWeight:X.sampleWeight};"ERROR"===zg?Id.level="ERROR_LEVEL_ERROR":"WARNING"===zg&&(Id.level="ERROR_LEVEL_WARNNING");var qj={isObfuscated:!0,browserStackInfo:vb},wb={pageUrl:window.location.href,kvPairs:[]};M("FEXP_EXPERIMENTS")&&(wb.experimentIds=M("FEXP_EXPERIMENTS"));var Jd=X.params;if(Jd)for(var Dg=u(Object.keys(Jd)),
Kd=Dg.next();!Kd.done;Kd=Dg.next()){var Eg=Kd.value;wb.kvPairs.push({key:"client."+Eg,value:String(Jd[Eg])})}var Fg=M("SERVER_NAME",void 0),Gg=M("SERVER_VERSION",void 0);Fg&&Gg&&(wb.kvPairs.push({key:"server.name",value:Fg}),wb.kvPairs.push({key:"server.version",value:Gg}));Mi("clientError",{errorMetadata:wb,stackTrace:qj,logMessage:Id});Yg()}}if(!P("suppress_error_204_logging")){var Ld,Fa=ia,xb=Fa.params||{},ua={urlParams:{a:"logerror",t:"jserror",type:Fa.name,msg:Fa.message.substr(0,250),line:Fa.lineNumber,
level:l,"client.name":xb.name},postParams:{url:M("PAGE_NAME",window.location.href),file:Fa.fileName},method:"POST"};xb.version&&(ua["client.version"]=xb.version);if(ua.postParams){Fa.stack&&(ua.postParams.stack=Fa.stack);for(var Hg=u(Object.keys(xb)),Md=Hg.next();!Md.done;Md=Hg.next()){var Ig=Md.value;ua.postParams["client."+Ig]=xb[Ig]}if(Ld=M("LATEST_ECATCHER_SERVICE_TRACKING_PARAMS",void 0))for(var Jg=u(Object.keys(Ld)),Nd=Jg.next();!Nd.done;Nd=Jg.next()){var Kg=Nd.value;ua.postParams[Kg]=Ld[Kg]}var Lg=
M("SERVER_NAME",void 0),Mg=M("SERVER_VERSION",void 0);Lg&&Mg&&(ua.postParams["server.name"]=Lg,ua.postParams["server.version"]=Mg)}sh(M("ECATCHER_REPORT_HOST","")+"/error_204",ua)}Wi.add(ia.message);Xi++}}}}}else throw Ua;}else console&&console.warn&&console.warn("The YouTube player is not attached to the DOM. API calls should be made after the onReady event. See more: https://developers.google.com/youtube/iframe_api_reference#Events")};function gj(a){return(0===a.search("cue")||0===a.search("load"))&&"loadModule"!==a}
function hj(a){return 0===a.search("get")||0===a.search("is")}
;function Z(a,b){if(!a)throw Error("YouTube player element ID required.");var c={title:"video player",videoId:"",width:640,height:360};if(b)for(var d in b)c[d]=b[d];Y.call(this,a,c,"player");this.C={};this.playerInfo={}}
x(Z,Y);r=Z.prototype;r.T=function(){var a=N(this.i,"playerVars");if(a){var b={},c;for(c in a)b[c]=a[c];a=b}else a={};window!=window.top&&document.referrer&&(a.widget_referrer=document.referrer.substring(0,256));if(c=N(this.i,"embedConfig")){if(D(c))try{c=JSON.stringify(c)}catch(d){console.error("Invalid embed config JSON",d)}a.embed_config=c}return a};
r.Y=function(a){var b=a.event;a=a.info;switch(b){case "apiInfoDelivery":if(D(a))for(var c in a)this.C[c]=a[c];break;case "infoDelivery":ij(this,a);break;case "initialDelivery":window.clearInterval(this.j);this.playerInfo={};this.C={};jj(this,a.apiInterface);ij(this,a);break;default:dj(this,b,a)}};
function ij(a,b){if(D(b))for(var c in b)a.playerInfo[c]=b[c]}
function jj(a,b){H(b,function(c){this[c]||("getCurrentTime"==c?this[c]=function(){var d=this.playerInfo.currentTime;if(1==this.playerInfo.playerState){var e=(Date.now()/1E3-this.playerInfo.currentTimeLastUpdated_)*this.playerInfo.playbackRate;0<e&&(d+=Math.min(e,1))}return d}:gj(c)?this[c]=function(){this.playerInfo={};
this.C={};fj(this,c,arguments);return this}:hj(c)?this[c]=function(){var d=0;
0===c.search("get")?d=3:0===c.search("is")&&(d=2);return this.playerInfo[c.charAt(d).toLowerCase()+c.substr(d+1)]}:this[c]=function(){fj(this,c,arguments);
return this})},a)}
r.getVideoEmbedCode=function(){var a=parseInt(N(this.i,"width"),10),b=parseInt(N(this.i,"height"),10),c=N(this.i,"host")+("/embed/"+N(this.i,"videoId"));zb.test(c)&&(-1!=c.indexOf("&")&&(c=c.replace(nb,"&amp;")),-1!=c.indexOf("<")&&(c=c.replace(ob,"&lt;")),-1!=c.indexOf(">")&&(c=c.replace(pb,"&gt;")),-1!=c.indexOf('"')&&(c=c.replace(qb,"&quot;")),-1!=c.indexOf("'")&&(c=c.replace(rb,"&#39;")),-1!=c.indexOf("\x00")&&(c=c.replace(yb,"&#0;")));return'<iframe width="'+a+'" height="'+b+'" src="'+c+'" frameborder="0" allow="accelerometer; autoplay; clipboard-write; encrypted-media; gyroscope; picture-in-picture" allowfullscreen></iframe>'};
r.getOptions=function(a){return this.C.namespaces?a?this.C[a].options||[]:this.C.namespaces||[]:[]};
r.getOption=function(a,b){if(this.C.namespaces&&a&&b)return this.C[a][b]};
function rj(a){if("iframe"!=a.tagName.toLowerCase()){var b=Rf(a,"videoid");b&&(b={videoId:b,width:Rf(a,"width"),height:Rf(a,"height")},new Z(a,b))}}
;E("YT.PlayerState.UNSTARTED",-1);E("YT.PlayerState.ENDED",0);E("YT.PlayerState.PLAYING",1);E("YT.PlayerState.PAUSED",2);E("YT.PlayerState.BUFFERING",3);E("YT.PlayerState.CUED",5);E("YT.get",function(a){return Nf[a]});
E("YT.scan",Qf);E("YT.subscribe",function(a,b,c){Ve.subscribe(a,b,c);Pf[a]=!0;for(var d in Nf)cj(Nf[d],a)});
E("YT.unsubscribe",function(a,b,c){Ue(a,b,c)});
E("YT.Player",Z);Y.prototype.destroy=Y.prototype.destroy;Y.prototype.setSize=Y.prototype.setSize;Y.prototype.getIframe=Y.prototype.pa;Y.prototype.addEventListener=Y.prototype.addEventListener;Z.prototype.getVideoEmbedCode=Z.prototype.getVideoEmbedCode;Z.prototype.getOptions=Z.prototype.getOptions;Z.prototype.getOption=Z.prototype.getOption;
Of.push(function(a){var b=a;b||(b=document);a=bb(b.getElementsByTagName("yt:player"));var c=b||document;if(c.querySelectorAll&&c.querySelector)b=c.querySelectorAll(".yt-player");else{var d;c=document;b=b||c;if(b.querySelectorAll&&b.querySelector)b=b.querySelectorAll(".yt-player");else if(b.getElementsByClassName){var e=b.getElementsByClassName("yt-player");b=e}else{e=b.getElementsByTagName("*");var f={};for(c=d=0;b=e[c];c++){var g=b.className,h;if(h="function"==typeof g.split)h=0<=Ta(g.split(/\s+/),
"yt-player");h&&(f[d++]=b)}f.length=d;b=f}}b=bb(b);H(ab(a,b),rj)});
"undefined"!=typeof YTConfig&&YTConfig.parsetags&&"onload"!=YTConfig.parsetags||Qf();var sj=B.onYTReady;sj&&sj();var tj=B.onYouTubeIframeAPIReady;tj&&tj();var uj=B.onYouTubePlayerAPIReady;uj&&uj();}).call(this);
