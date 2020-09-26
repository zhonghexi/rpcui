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
Togglebar('tablist','tabcon','current')


function Togglebar(tab,content,current){
    var tabs = document.getElementById(tab);
    var lis = tabs.children;
    var contents = document.getElementById(content);
    var item = contents.children;
    for (var j = 0; j < lis.length; j++) {
        lis[j].index = j
        lis[j].onclick = function(){
            for (var j = 0; j <lis.length; j++) {
                lis[j].className = ''
                item[j].style.display = 'none'
            }
            this.className = current
            item[this.index].style.display = 'block'
        }
    }

}