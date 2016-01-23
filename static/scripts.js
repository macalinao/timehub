$(function () {
  setTimeout(function() {
    $(window).scrollTop(10000)
  }, 0)
  $(window).scroll(function (e) {
    console.log($(this).scrollTop())
  })
})
