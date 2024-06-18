/**
* @returns {void}
*/
function run() {

  if (!config.video.url) {
    console.error('No video URL provided');
    return;
  }

  preload(config.video.url);

  const initialScrollPosition = window.scrollY;

  function handleScroll() {
    let scrollPosition = window.scrollY;
    if (Math.abs(initialScrollPosition - scrollPosition) > 100 && config.video.small.height > 0) {
      setupIntroVideo({ bubble, cta });
      window.removeEventListener('scroll', handleScroll);
    }
  }

  window.addEventListener('scroll', handleScroll);
}

run();

