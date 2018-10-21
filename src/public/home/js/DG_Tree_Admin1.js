$(function(){
    if($("#Tree_Contianer").length>0&&$("#InitTreeData").length>0){

        var view_lower=$('#view_lower').val();
        var view_Uplower = $("#view_Uplower").val() || "";
        var is_show=$('#is_show').val();
        var K = 0 ;

        function _CreatTrHtml(data,level,Pid,K,index){
            var _Width = level * 24;
            var strVar = "";
            strVar += "<tr  parentid='"+Pid+"' childId='"+ data.user.id +"' state='Down'>";
            /*strVar += "	<td class=\"table-id\">"+index+"<\/td>";*/
            strVar += "	<td class=\"table-title\">";
            strVar += "		<span style='width:"+_Width+"px;display:inline-block;'></span>" ;
            if(level>0){
                strVar += "		<span  class='folder-line'></span>"  ;
            }
            if(is_show==1){
                strVar += "		<span  class='look_child folder-open' ></span>" ;
            }else{
                strVar += "		<span  class='look_child folder-open' level='"+level+"' pid='"+index+"' guid='"+data.user.id+"'></span>" ;
            }

            strVar += 	"用户(" + data.user.account + " )  手机("+ data.user.mobile + " )  团队购买矿机的总量:"+ data.user_tree.team_buy_num + "   团队用户量:"+ data.user_tree.team_num + " ";
            strVar += "<\/td> ";
            if(is_show==1){
                strVar += "	<td class=\"table-set\" style='text-align: center;'><a href='javascript:void(0)' level='"+level+"' pid='"+index+"' class='look_child' guid='"+data.user.id+"'></a>";
            }else{
                strVar += "	<td class=\"table-set\" style='text-align: center;'><a href='javascript:void(0)' level='"+level+"' pid='"+index+"' class='look_child' guid='"+data.user.id+"'>"+view_lower+"</a>";
            }

            if(view_Uplower!=""){
                strVar += "	&nbsp;&nbsp;<a href='javascript:void(0)' level='" + level + "' pid='" + index + "' class='look_parent' guid='" + data.user_tree.direct_id + "'>" + view_Uplower + "</a>";
            }
            strVar += "<\/td>";
            strVar += "<\/tr>";
            return strVar;
        }
        var _theInitData = $("#InitTreeData").val();
        if(_theInitData!=""){
            var _theInitDataJson = $.parseJSON(_theInitData);
            if(_theInitDataJson!=undefined&&_theInitDataJson.length>0){
                for(var i in _theInitDataJson){
                    var Index = Number(i)+1;
                    var _theTrStr = _CreatTrHtml(_theInitDataJson[i],0,0,K,Index);
                    if(_theTrStr!=""){
                        $("#Tree_Contianer").append($(_theTrStr));
                    }
                }
            }
        }


        var sk=true;
        //绑定点击事件 ---- 查看下级
        $("#Tree_Contianer").on("click",".look_child",function(){

            var _TheLevel = $(this).attr("level");
            if($(this).parents("tr").eq(0).hasClass("curr")){
                var _this = $(this);
                //_this.parents("tr").removeClass("curr");

                var _ChildId = _this.parents("tr").attr("childid") == undefined ? "" : $(this).attr("childid");
                var _parentObj = _this.parents("tr").eq(0);
                Menu_contraction("Tree_Contianer", _ChildId, _parentObj);

            }else{

                var _this = $(this);
                if($("#GetChildData").length>0&&$("#GetChildData").val()!=""){
                    var _theId = _this.attr("guid") == undefined ? 0 : _this.attr("guid");
                    /* if(sk==_theId){
                         return false;
                     }
                     sk=_theId;*/
                    if($("#Tree_Contianer tr[parentid='"+_theId+"']").length>0){
                        var _ChildId = _this.parents("tr").attr("childid") == undefined ? "" : $(this).attr("childid");
                        var _parentObj = _this.parents("tr").eq(0);
                        Menu_contraction("Tree_Contianer", _ChildId, _parentObj);
                    }else{
                        var _theUrl = $("#GetChildData").val();
                        $.ajax({
                            type:"POST",
                            url:_theUrl,
                            data:{
                                id:_theId
                            },
                            dataType:"json",
                            success:function(data){
                                if(data){
                                    K = K+1;
                                    _this.parents("tr").addClass("curr");
                                    _this.parents("tr").attr("state","Up");
                                    if(data.data_children.length>0){
                                        for(var i in data.data_children){
                                            var _ChildLevel = Number(_TheLevel)+1;
                                            var _theString = _CreatTrHtml(data.data_children[i],_ChildLevel,_theId,K,"");
                                            _this.parents("tr").after($(_theString));
                                        }
                                    }
                                }
                            }
                        });
                    }
                }
            }
        });

        //绑定点击事件----查看上级
        $("#Tree_Contianer").on("click",".look_parent",function(event){
            var _theParentId = $(this).parents("tr").eq(0).attr("parentid") || "";
            event = event || window.event;
            event.stopPropagation();
            if(_theParentId!=""&&_theParentId!=0){
                var _theLengths = $("#Tree_Contianer").find("tr[childid='"+_theParentId+"']").length;
                if(_theLengths > 0){
                    $("#Tree_Contianer").find("tr[childid='"+_theParentId+"']").find(".look_child").eq(0).trigger("click");
                }
            }else if(_theParentId==0){
                var _theUrl = $("#GetParentData").val() || "";
                var _theGuid = $(this).attr("guid") || "";
                if(_theUrl!=""&&_theGuid!=0){
                    window.location.href=_theUrl+"/"+_theGuid;
                   // window.location.href
                }
            }
        });
    }


//======================lrxiang无限极菜单收缩======================//

//递归菜单收缩或者展开
    function Menu_contraction(ContObj, child_id, thePObj) {   //id ， 0 ，undefined ，undefined

        if (child_id == undefined) { child_id = 0; }
        if (ContObj == undefined) { ContObj = 'body'; }

        if (thePObj == undefined) {

            $("#" + ContObj + " tr[parentid='" + child_id + "']").show();
            $("#" + ContObj + " tr[parentid='" + child_id + "']").attr("state","Down");

        } else {

            var This_SlideType = thePObj.attr("state") == undefined ? "Down" : thePObj.attr("state") ;
            if (This_SlideType == "Down") {     //展开所有子集
                var _ChildId = thePObj.attr("childid") == undefined ? "" : thePObj.attr("childid");
                if (_ChildId != "") {

                    $("#" + ContObj + " tr[parentid='" + _ChildId + "']").show();
                    $("#" + ContObj + " tr[parentid='" + _ChildId + "']").attr("state", "Down");
                    thePObj.attr("state", "Up");
                }

            } else {    //收缩所有的子集
                var _ChildId = thePObj.attr("childid") == undefined ? "" : thePObj.attr("childid");
                if (_ChildId != "") {
                    $("#" + ContObj + " tr[parentid='" + _ChildId + "']").each(function () {
                        var This_Type = $(this).attr("state") == undefined ? "Down" : $(this).attr("state");
                        if (This_Type == "Down") { //当前子类的后代属于隐藏状态
                            $(this).hide();
                        } else {
                            $(this).attr("state","Up")
                            //把所有子集隐藏
                            var _CChildId = $(this).attr("childid") == undefined ? "" : $(this).attr("childid");
                            if (_CChildId != "") {
                                Menu_contraction(ContObj, _ChildId, $(this));
                                $(this).hide();
                            }
                        }
                    });
                    thePObj.attr("state", "Down");
                    //thePObj.hide();
                } else {
                    thePObj.attr("state", "Down");
                }
            }
        }
    }


//======================lrxiang无限极菜单收缩======================//


});