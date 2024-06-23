/**
* @returns {void}
*/
function run() {

  if (!videoConfig.url) {
    console.error('No video URL provided');
    return;
  }

  preload(videoConfig.url);

  const initialScrollPosition = window.scrollY;

  function handleScroll() {
    let scrollPosition = window.scrollY;
    if (Math.abs(initialScrollPosition - scrollPosition) > 100 && videoConfig.small.height > 0) {
      setupIntroVideo({ bubble, cta });
      window.removeEventListener('scroll', handleScroll);
    }
  }

  window.addEventListener('scroll', handleScroll);
}

run();

