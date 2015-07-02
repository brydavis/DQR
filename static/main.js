$(document).ready(function (e) {
	$(".qrate").each(function (i) {
		if (parseFloat($(this).html()) < 90) {
			console.log(parseFloat($(this).html()))
			$(this).toggleClass("danger")
		} else {
			$(this).toggleClass("info")

		}
	})
})

