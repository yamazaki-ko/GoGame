$(function() {
	var showFlag = false;
	var topBtn = $('#pagetopbtn');    
	topBtn.css('bottom', '-100px');
	var showFlag = false;
	$(window).scroll(function () {
		if ($(this).scrollTop() > 100) {
			if (showFlag == false) {
				showFlag = true;
				topBtn.stop().animate({'bottom' : '15px'}, 200); 
			}
		} else {
			if (showFlag) {
				showFlag = false;
				topBtn.stop().animate({'bottom' : '-100px'}, 200); 
			}
		}
	});
	topBtn.click(function () {
		$('body,html').animate({
			scrollTop: 0
		}, 500);
		return false;
	});
	$("img.lazy").lazyload();
});