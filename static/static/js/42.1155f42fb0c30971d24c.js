webpackJsonp([42],{511:function(t,e,a){var n=a(6)(a(599),a(708),null,null,null);t.exports=n.exports},599:function(t,e,a){"use strict";Object.defineProperty(e,"__esModule",{value:!0}),e.default={data:function(){return{movieList:[{name:"肖申克的救赎",url:"https://movie.douban.com/subject/1292052/",rate:9.6},{name:"这个杀手不太冷",url:"https://movie.douban.com/subject/1295644/",rate:9.4},{name:"霸王别姬",url:"https://movie.douban.com/subject/1291546/",rate:9.5},{name:"阿甘正传",url:"https://movie.douban.com/subject/1292720/",rate:9.4},{name:"美丽人生",url:"https://movie.douban.com/subject/1292063/",rate:9.5},{name:"千与千寻",url:"https://movie.douban.com/subject/1291561/",rate:9.2},{name:"辛德勒的名单",url:"https://movie.douban.com/subject/1295124/",rate:9.4},{name:"海上钢琴师",url:"https://movie.douban.com/subject/1292001/",rate:9.2},{name:"机器人总动员",url:"https://movie.douban.com/subject/2131459/",rate:9.3},{name:"盗梦空间",url:"https://movie.douban.com/subject/3541415/",rate:9.2}],randomMovieList:[]}},methods:{changeLimit:function(){this.randomMovieList=function(t,e){var a=[];for(var n in t)a.push(t[n]);for(var o=[],r=0;r<e&&a.length>0;r++){var s=Math.floor(Math.random()*a.length);o[r]=a[s],a.splice(s,1)}return o}(this.movieList,5)}},mounted:function(){this.changeLimit()}}},708:function(t,e){t.exports={render:function(){var t=this,e=t.$createElement,a=t._self._c||e;return a("div",{staticClass:"animated fadeIn"},[a("Row",[a("Col",{attrs:{span:"24"}},[a("div",{staticClass:"doc-header"},[a("Card",{staticStyle:{width:"350px"}},[a("p",{slot:"title"},[a("Icon",{attrs:{type:"ios-film-outline"}}),t._v("\n            经典电影\n        ")],1),t._v(" "),a("a",{attrs:{href:"#"},on:{click:function(e){e.preventDefault(),t.changeLimit(e)}},slot:"extra"},[a("Icon",{attrs:{type:"ios-loop-strong"}}),t._v("\n            换一换\n        ")],1),t._v(" "),a("ul",t._l(t.randomMovieList,function(e){return a("li",[a("a",{attrs:{href:e.url,target:"_blank"}},[t._v(t._s(e.name))]),t._v(" "),a("span",[t._l(4,function(t){return a("Icon",{key:t,attrs:{type:"ios-star",color:"#ffac2d"}})}),e.rate>=9.5?a("Icon",{attrs:{type:"ios-star",color:"#ffac2d"}}):a("Icon",{attrs:{type:"ios-star-half",color:"#ffac2d"}}),t._v("\n                    "+t._s(e.rate)+"\n                ")],2)])}))])],1),t._v(" "),a("div",{staticClass:"doc-content"},[a("h5",[t._v("基本用法")]),t._v(" "),a("p",[t._v("自定义标题、额外操作和主体内容，可以完全自由控制各个部分，也可以结合其它组件一起使用，较为灵活。")])])])],1)],1)},staticRenderFns:[]}}});
//# sourceMappingURL=42.1155f42fb0c30971d24c.js.map