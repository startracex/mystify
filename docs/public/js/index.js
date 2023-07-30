const postsImagesApi = '/img';
const img = $(".markdown img");
if (img.length)
    img.each(function () {
        if ($(this).attr('src').indexOf("/public") != 0)
            $(this).attr('src', `${postsImagesApi}?pathname=${location.pathname}&&filepath=${$(this).attr('src')}`)
    })

var list = [];
$(function () {
    $('.markdown h1,.markdown h2,.markdown h3,.markdown h4,.markdown h5,.markdown h6').each(function () {
        $(this).attr('id', $(this).text());
        list.push(`<a class=` + `${$(this)[0].nodeName}` + ` href="#` + $(this).text() + `" offset="` + $(this).offset().top + `">` + $(this).text() + `</a>`)
    });
    $('#-overview').prepend(list);
    $('#-m-overview').prepend(list);
})
const tocopy = '<svg width="20" height="20" viewBox="0 0 48 48" fill="none"><path d="M13 12.4316V7.8125C13 6.2592 14.2592 5 15.8125 5H40.1875C41.7408 5 43 6.2592 43 7.8125V32.1875C43 33.7408 41.7408 35 40.1875 35H35.5163" stroke="#333" stroke-width="3" stroke-linecap="round" stroke-linejoin="round"/><path d="M32.1875 13H7.8125C6.2592 13 5 14.2592 5 15.8125V40.1875C5 41.7408 6.2592 43 7.8125 43H32.1875C33.7408 43 35 41.7408 35 40.1875V15.8125C35 14.2592 33.7408 13 32.1875 13Z" fill="none" stroke="#333" stroke-width="3" stroke-linejoin="round"/></svg>'
const ifcopy = '<svg width="20" height="20" viewBox="0 0 48 48" fill="none"><path d="M10 24L20 34L40 14" stroke="#333" stroke-width="3" stroke-linecap="round" stroke-linejoin="round"/></svg>'
const delaytime = 750
$(function () {
    var pres = $("pre");
    if (pres.length)
        pres.each(function () {
            var tx = $(this).children("code").text();
            var btn = $('<span class="copy">' + tocopy + '</span>').attr("data-clipboard-text", tx);
            $(this).prepend(btn);
            var c = new ClipboardJS(btn[0]);
            c.on("success", function () {
                btn.addClass("Copied").html(ifcopy);
                setTimeout(function () {
                    btn.html(tocopy).removeClass("Copied");
                }, delaytime);
            });
            c.on("error", function () {
                btn.html('<svg width="20" height="20" viewBox="0 0 48 48" fill="none"><path d="M14 14L34 34" stroke="#333" stroke-width="3" stroke-linecap="round" stroke-linejoin="round"/><path d="M14 34L34 14" stroke="#333" stroke-width="3" stroke-linecap="round" stroke-linejoin="round"/></svg>');
            });
        });
});
var $box = $("#-overview-x");
eleDragStatus($box)
function eleDragStatus($box) {
    $box.mousedown(function (event) {
        var old_left = event.pageX;
        var old_top = event.pageY;
        var old_position_left = $(this).position().left;
        var old_position_top = $(this).position().top;
        $(document).mousemove(function (event) {
            var new_left = event.pageX;
            var new_top = event.pageY;
            var chang_x = new_left - old_left;
            var change_y = new_top - old_top;
            var new_position_top = old_position_top + change_y;
            if (new_position_top < 0) new_position_top = 0;
            if (new_position_top > document.body.clientHeight - $box.height()) new_position_top = document.body.clientHeight - $box.height();
            $box.css({
                top: new_position_top + 'px',
            })
        });
        $box.mouseup(function () {
            $(document).off("mousemove");
        });
    });
}
$(function () {
    let scro = 0;
    scroll()
    window.addEventListener('scroll', () => {
        scroll()
    })
    function scroll() {
        const scrollTop = document.documentElement.scrollTop || document.body.scrollTop;
        if (scrollTop > scro) {
            document.querySelector("#-overview-x").style.display = "block";
        }
        scro = scrollTop;
        document.querySelectorAll('#-overview a').forEach((a) => {
            if (scrollTop + document.body.clientHeight / 11 < a.getAttribute('offset') && a.getAttribute('offset') < scrollTop + document.body.clientHeight / 1.1) {
                a.classList.add('reading')
            } else {
                a.classList.remove('reading')
            }
        });
    }
})
