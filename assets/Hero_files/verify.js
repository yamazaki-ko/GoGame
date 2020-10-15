setTimeout(function(){
    'use strict';
    try{
        var stringifyFunc = null;
		if(window.JSON){
			stringifyFunc = window.JSON.stringify;
		} else {
			if(window.parent && window.parent.JSON){
				stringifyFunc = window.parent.JSON.stringify;
			}
		}
		if(!stringifyFunc){
			return;
		}
        var msg = {
            action: 'notifyBrandShieldAdEntityInformation',
            bsAdEntityInformation: {
                brandShieldId:'94d15e6d875446f9984846263b59dbe3',
                comparisonItems:[{name : 'cmp', value : 10257004},{name : 'plmt', value : 16826691}]
            }
        };
        var msgString = stringifyFunc(msg);
        var bst2tWin = null;

        var findAndSendMessage = function() {
            if (!bst2tWin) {
                var frame = document.getElementById('bst2t_204029124176');
                if (frame) {
                    bst2tWin = frame.contentWindow;
                }
            }

            if (bst2tWin) {
                bst2tWin.postMessage(msgString, '*');
            }
        };

        findAndSendMessage();
        setTimeout(findAndSendMessage, 50);
        setTimeout(findAndSendMessage, 500);
    } catch(err){}
}, 10);var dvObj = $dvbs;function np764531(g,i){function d(){function d(){function f(b,a){b=(i?"dvp_":"")+b;e[b]=a}var e={},a=function(b){for(var a=[],c=0;c<b.length;c+=2)a.push(String.fromCharCode(parseInt(b.charAt(c)+b.charAt(c+1),32)));return a.join("")},h=window[a("3e313m3937313k3f3i")];h&&(a=h[a("3g3c313k363f3i3d")],f("pltfrm",a));(function(){var a=e;e={};if (a['pltfrm'])dvObj.registerEventCall(g,a,2E3,true)})()}try{d()}catch(f){}}try{dvObj.pubSub.subscribe(dvObj==window.$dv?"ImpressionServed":"BeforeDecisionRender",g,"np764531",d)}catch(f){}}
;np764531("94d15e6d875446f9984846263b59dbe3",false);


try{__tagObject_callback_204029124176({ImpressionID:"94d15e6d875446f9984846263b59dbe3", ServerPublicDns:"tps802.doubleverify.com"});}catch(e){}
try{$dvbs.pubSub.publish('BeforeDecisionRender', "94d15e6d875446f9984846263b59dbe3");}catch(e){}
try{__verify_callback_204029124176({
ResultID:2,
Passback:"",
AdWidth:300,
AdHeight:600});}catch(e){}
try{$dvbs.pubSub.publish('AfterDecisionRender', "94d15e6d875446f9984846263b59dbe3");}catch(e){}
