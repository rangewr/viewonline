webpackJsonp([34],{489:function(t,a,e){function o(t){e(658)}var i=e(6)(e(577),e(702),o,null,null);t.exports=i.exports},577:function(t,a,e){"use strict";Object.defineProperty(a,"__esModule",{value:!0}),a.default={name:"buttons",data:function(){return{loading:!1,loading2:!1,now_show_panel:"1",sampleData1:{medications:[{aceInhibitors:[{name:"lisinopril",strength:"10 mg Tab",dose:"1 tab",route:"PO",sig:"daily",pillCount:"#90",refills:"Refill 3"}],antianginal:[{name:"nitroglycerin",strength:"0.4 mg Sublingual Tab",dose:"1 tab",route:"SL",sig:"q15min PRN",pillCount:"#30",refills:"Refill 1"}],anticoagulants:[{name:"warfarin sodium",strength:"3 mg Tab",dose:"1 tab",route:"PO",sig:"daily",pillCount:"#90",refills:"Refill 3"}],betaBlocker:[{name:"metoprolol tartrate",strength:"25 mg Tab",dose:"1 tab",route:"PO",sig:"daily",pillCount:"#90",refills:"Refill 3"}],diuretic:[{name:"furosemide",strength:"40 mg Tab",dose:"1 tab",route:"PO",sig:"daily",pillCount:"#90",refills:"Refill 3"}],mineral:[{name:"potassium chloride ER",strength:"10 mEq Tab",dose:"1 tab",route:"PO",sig:"daily",pillCount:"#90",refills:"Refill 3"}]}],labs:[{name:"Arterial Blood Gas",time:"Today",location:"Main Hospital Lab"},{name:"BMP",time:"Today",location:"Primary Care Clinic"},{name:"BNP",time:"3 Weeks",location:"Primary Care Clinic"},{name:"BUN",time:"1 Year",location:"Primary Care Clinic"},{name:"Cardiac Enzymes",time:"Today",location:"Primary Care Clinic"},{name:"CBC",time:"1 Year",location:"Primary Care Clinic"},{name:"Creatinine",time:"1 Year",location:"Main Hospital Lab"},{name:"Electrolyte Panel",time:"1 Year",location:"Primary Care Clinic"},{name:"Glucose",time:"1 Year",location:"Main Hospital Lab"},{name:"PT/INR",time:"3 Weeks",location:"Primary Care Clinic"},{name:"PTT",time:"3 Weeks",location:"Coumadin Clinic"},{name:"TSH",time:"1 Year",location:"Primary Care Clinic"}],imaging:[{name:"Chest X-Ray",time:"Today",location:"Main Hospital Radiology"},{name:"Chest X-Ray",time:"Today",location:"Main Hospital Radiology"},{name:"Chest X-Ray",time:"Today",location:"Main Hospital Radiology"}]}}}}},625:function(t,a,e){a=t.exports=e(477)(!0),a.push([t.i,"","",{version:3,sources:[],names:[],mappings:"",file:"JsonTree.vue",sourceRoot:""}])},658:function(t,a,e){var o=e(625);"string"==typeof o&&(o=[[t.i,o,""]]),o.locals&&(t.exports=o.locals);e(478)("087cd678",o,!0)},702:function(t,a){t.exports={render:function(){var t=this,a=t.$createElement,e=t._self._c||a;return e("div",{staticClass:"animated fadeIn"},[e("Row",[e("Col",{attrs:{xs:24,sm:24,md:12,lg:12}},[e("div",{staticClass:"doc-header"},[e("Collapse",{attrs:{accordion:""},model:{value:t.now_show_panel,callback:function(a){t.now_show_panel=a},expression:"now_show_panel"}},[e("Panel",{attrs:{name:"1"}},[t._v("\n                      maxDepth:1\n                      "),e("p",{slot:"content"},[e("tree-view",{attrs:{data:t.sampleData1,options:{maxDepth:1}}})],1)]),t._v(" "),e("Panel",{attrs:{name:"2"}},[t._v("\n                         maxDepth:2\n                      "),e("p",{slot:"content"},[e("tree-view",{attrs:{data:t.sampleData1,options:{maxDepth:2}}})],1)]),t._v(" "),e("Panel",{attrs:{name:"3"}},[t._v("\n                      maxDepth:3\n                      "),e("p",{slot:"content"},[e("tree-view",{attrs:{data:t.sampleData1,options:{maxDepth:3}}})],1)])],1)],1),t._v(" "),e("div",{staticClass:"doc-content"},[e("h5",[t._v("JSON展示列表")]),t._v(" "),e("p",[t._v("配合Collapse 折叠面板很轻易的就可以展示又臭又长的json字符串")])])]),t._v(" "),e("Col",{attrs:{xs:24,sm:24,md:12,lg:12}},[e("div",{staticClass:"highlight",staticStyle:{background:"#f0f0f0"}},[e("pre",{staticStyle:{"line-height":"125%"}},[e("span",{staticStyle:{color:"#062873","font-weight":"bold"}},[t._v("<template>")]),t._v("\n\n       "),e("span",{staticStyle:{color:"#062873","font-weight":"bold"}},[t._v("<Collapse")]),t._v(" "),e("span",{staticStyle:{color:"#4070a0"}},[t._v('v-model="now_show_panel"')]),t._v(" "),e("span",[t._v("accordion")]),e("span",{staticStyle:{color:"#062873","font-weight":"bold"}},[t._v(">")]),t._v(" \n\n         "),e("span",{staticStyle:{color:"#062873","font-weight":"bold"}},[t._v("<Panel")]),t._v(" "),e("span",{staticStyle:{color:"#4070a0"}},[t._v('name="1"')]),e("span",{staticStyle:{color:"#062873","font-weight":"bold"}},[t._v(">")]),t._v("\n              maxDepth:1\n              "),e("span",{staticStyle:{color:"#062873","font-weight":"bold"}},[t._v("<p")]),t._v(" "),e("span",{staticStyle:{color:"#4070a0"}},[t._v('slot="content"')]),e("span",{staticStyle:{color:"#062873","font-weight":"bold"}},[t._v(">")]),t._v("\n                "),e("span",{staticStyle:{color:"#062873","font-weight":"bold"}},[t._v("<tree-view")]),t._v(" "),e("span",{staticStyle:{color:"#4070a0"}},[t._v(':data="sampleData1"')]),t._v(" "),e("span",{staticStyle:{color:"#4070a0"}},[t._v(':options="{maxDepth: 1}"')]),e("span",{staticStyle:{color:"#062873","font-weight":"bold"}},[t._v(">\n                </tree-view>")]),t._v("\n              "),e("span",{staticStyle:{color:"#062873","font-weight":"bold"}},[t._v("</p>")]),t._v("\n          "),e("span",{staticStyle:{color:"#062873","font-weight":"bold"}},[t._v("</Panel>")]),t._v("\n          "),e("span",{staticStyle:{color:"#062873","font-weight":"bold"}},[t._v("<Panel")]),t._v(" "),e("span",{staticStyle:{color:"#4070a0"}},[t._v('name="2"')]),t._v(" "),e("span",{staticStyle:{color:"#062873","font-weight":"bold"}},[t._v(">")]),t._v("\n                 maxDepth:2\n              "),e("span",{staticStyle:{color:"#062873","font-weight":"bold"}},[t._v("<p")]),t._v(" "),e("span",{staticStyle:{color:"#4070a0"}},[t._v('slot="content"')]),e("span",{staticStyle:{color:"#062873","font-weight":"bold"}},[t._v(">")]),t._v("\n                "),e("span",{staticStyle:{color:"#062873","font-weight":"bold"}},[t._v("<tree-view")]),t._v(" "),e("span",{staticStyle:{color:"#4070a0"}},[t._v(':data="sampleData1"')]),t._v(" "),e("span",{staticStyle:{color:"#4070a0"}},[t._v(':options="{maxDepth: 2}"')]),e("span",{staticStyle:{color:"#062873","font-weight":"bold"}},[t._v(">\n                </tree-view>")]),t._v("\n              "),e("span",{staticStyle:{color:"#062873","font-weight":"bold"}},[t._v("</p>")]),t._v("\n          "),e("span",{staticStyle:{color:"#062873","font-weight":"bold"}},[t._v("</Panel>")]),t._v("\n          "),e("span",{staticStyle:{color:"#062873","font-weight":"bold"}},[t._v("<Panel")]),t._v(" "),e("span",{staticStyle:{color:"#4070a0"}},[t._v('name="3"')]),e("span",{staticStyle:{color:"#062873","font-weight":"bold"}},[t._v(">")]),t._v("\n              maxDepth:3\n              "),e("span",{staticStyle:{color:"#062873","font-weight":"bold"}},[t._v("<p")]),t._v(" "),e("span",{staticStyle:{color:"#4070a0"}},[t._v('slot="content"')]),e("span",{staticStyle:{color:"#062873","font-weight":"bold"}},[t._v(">")]),t._v("\n                "),e("span",{staticStyle:{color:"#062873","font-weight":"bold"}},[t._v("<tree-view")]),t._v(" "),e("span",{staticStyle:{color:"#4070a0"}},[t._v(':data="sampleData1"')]),t._v(" "),e("span",{staticStyle:{color:"#4070a0"}},[t._v(':options="{maxDepth: 3}"')]),e("span",{staticStyle:{color:"#062873","font-weight":"bold"}},[t._v(">\n                </tree-view>")]),t._v("\n              "),e("span",{staticStyle:{color:"#062873","font-weight":"bold"}},[t._v("</p>")]),t._v("\n          "),e("span",{staticStyle:{color:"#062873","font-weight":"bold"}},[t._v("</Panel>")]),t._v("\n             \n        "),e("span",{staticStyle:{color:"#062873","font-weight":"bold"}},[t._v("</Collapse>")]),t._v("\n\n"),e("span",{staticStyle:{color:"#062873","font-weight":"bold"}},[t._v("</template>")]),t._v("\n"),e("span",{staticStyle:{color:"#062873","font-weight":"bold"}},[t._v("<script>")]),t._v('\n\nexport default {\n  name: \'buttons\',\n\n   data () {\n            return {\n                value:1,\n                 sampleData1: {"medications":[{"aceInhibitors":[{"name":"lisinopril","strength":"10 mg Tab","dose":"1 tab","route":"PO","sig":"daily","pillCount":"#90","refills":"Refill 3"}],"antianginal":[{"name":"nitroglycerin","strength":"0.4 mg Sublingual Tab","dose":"1 tab","route":"SL","sig":"q15min PRN","pillCount":"#30","refills":"Refill 1"}],"anticoagulants":[{"name":"warfarin sodium","strength":"3 mg Tab","dose":"1 tab","route":"PO","sig":"daily","pillCount":"#90","refills":"Refill 3"}],"betaBlocker":[{"name":"metoprolol tartrate","strength":"25 mg Tab","dose":"1 tab","route":"PO","sig":"daily","pillCount":"#90","refills":"Refill 3"}],"diuretic":[{"name":"furosemide","strength":"40 mg Tab","dose":"1 tab","route":"PO","sig":"daily","pillCount":"#90","refills":"Refill 3"}],"mineral":[{"name":"potassium chloride ER","strength":"10 mEq Tab","dose":"1 tab","route":"PO","sig":"daily","pillCount":"#90","refills":"Refill 3"}]}],"labs":[{"name":"Arterial Blood Gas","time":"Today","location":"Main Hospital Lab"},{"name":"BMP","time":"Today","location":"Primary Care Clinic"},{"name":"BNP","time":"3 Weeks","location":"Primary Care Clinic"},{"name":"BUN","time":"1 Year","location":"Primary Care Clinic"},{"name":"Cardiac Enzymes","time":"Today","location":"Primary Care Clinic"},{"name":"CBC","time":"1 Year","location":"Primary Care Clinic"},{"name":"Creatinine","time":"1 Year","location":"Main Hospital Lab"},{"name":"Electrolyte Panel","time":"1 Year","location":"Primary Care Clinic"},{"name":"Glucose","time":"1 Year","location":"Main Hospital Lab"},{"name":"PT/INR","time":"3 Weeks","location":"Primary Care Clinic"},{"name":"PTT","time":"3 Weeks","location":"Coumadin Clinic"},{"name":"TSH","time":"1 Year","location":"Primary Care Clinic"}],"imaging":[{"name":"Chest X-Ray","time":"Today","location":"Main Hospital Radiology"},{"name":"Chest X-Ray","time":"Today","location":"Main Hospital Radiology"},{"name":"Chest X-Ray","time":"Today","location":"Main Hospital Radiology"}]}\n            }\n        },\n \n}\n\n\n'),e("span",{staticStyle:{color:"#062873","font-weight":"bold"}},[t._v("<\/script>")]),t._v("\n")])])])],1)],1)},staticRenderFns:[]}}});
//# sourceMappingURL=34.e38b0017c98a699a0636.js.map