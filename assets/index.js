// var ul = document.getElementById('tablist')
// var lis = ul.children
// var tabcon = document.getElementById('tabcon')
// var items = tabcon.children
// for (var j = 0; j < lis.length; j++) {
//     lis[j].index = j
//     lis[j].onclick = function(){
//         for (var j = 0; j <lis.length; j++) {
//             lis[j].className = ''
//             items[j].style.display = 'none'
//         }
//         this.className = 'current'
//         items[this.index].style.display = 'block'
//     }
// }


// var lists = document.querySelector('.item-list')
// var lis1 = lists.children
// console.log(lis1);
// for (let j = 0; j < lis1.length; j++) {
//     // lis[j].index = j
//    lis1[j].onclick = function(){
//        for (var j = 0; j < lis1.length; j++) {
//           lis1[j].className = ''
//        }
//        this.className = 'current'
//    }
// }
//   切换栏   切换内容   样式
// Togglebar('tablist','tabcon','current')
Togglebar({
    tabList:'tab-list',
    tabContent:'tab-con',
    current:'current'
});

Togglebar({
    tabList:'item-list',
    tabContent:null,
    current:'current'
});




function Togglebar(options){
    var tabs = document.getElementById(options.tabList);
    var lis = tabs.children;
    if(!(options.tabContent == null)){
        var contents = document.getElementById(options.tabContent);
        var item = contents.children;
        var flag = 1
    };
    for (var j = 0; j < lis.length; j++) {
        lis[j].index = j
        lis[j].onclick = function(){
            for (var j = 0; j <lis.length; j++) {
                lis[j].className = ''
                if(flag){
                    item[j].style.display = 'none'
                }
            }
            this.className = options.current
            if(flag){
                item[this.index].style.display = 'block'
            }
        }
    }

}