/**
* @param {number} area
* @param {number} aspectRatio
* returns {number}
*/
function calculateWidth(area, aspectRatio) {
  return Math.sqrt(area * aspectRatio);
}

function loadContainer() {
  if (container) {
    return container
  }

  let unableToFindContainer = false;

  if (config.target !== null) {
    try {
      container = document.querySelector(config.target);
      if (container === null) {
        unableToFindContainer = true;
      } else {
        return container;
      }
    } catch (e) {
      unableToFindContainer = true;
    }
  }

  if (config.target === null || unableToFindContainer) {
    const body = document.querySelector('body');
    container = document.createElement('div');

    body.appendChild(container);
  }

  return container;
}

var container = loadContainer();
var video = null;

/**
* @param {string} videoUrl
* @returns {void}
*/
function preload(videoUrl) {
  video = document.createElement('video');

  video.addEventListener('loadeddata', () => {
    const ratio = video.videoWidth / video.videoHeight;
    config.video.small.width = calculateWidth(284 * 160, ratio);
    config.video.small.height = config.video.small.width / ratio;

    config.video.large.width = calculateWidth(480 * 270, ratio);
    config.video.large.height = config.video.large.width / ratio;
  });

  video.classList.add('iv-player');

  video.muted = true;
  video.loop = true;
  video.draggable = false;
  video.src = videoUrl;

  container.appendChild(video);
}



/**
* @param {HTMLDivElement} bubble
* @param {HTMLDivElement} cta
* @returns {void}
*/
function setupIntroVideo({ bubble, cta }) {
  const card = document.createElement('div');
  card.classList.add('iv-card');

  card.style.width = `${config.video.small.width}px`;
  card.style.height = `${config.video.small.height}px`;


  const videoWrapper = document.createElement('div');
  videoWrapper.classList.add('iv-player-wrapper');

  /**
  * @type {HTMLVideoElement}
  */
  video.style.display = 'block';

  const progressBar = document.createElement('progress');
  progressBar.classList.add('iv-progressbar');
  progressBar.value = 0;
  progressBar.max = 100;

  const button = document.createElement('button');
  button.classList.add('iv-close-button');
  button.innerHTML = '&times;';

  video.addEventListener('timeupdate', function() {
    const percentage = (video.currentTime / video.duration) * 100;
    progressBar.value = percentage;
  });

  button.onclick = () => {
    card.style.opacity = 0;
    setTimeout(() => {
      card.remove();
    }, 500);
  }

  videoWrapper.onclick = () => {
    card.style.height = `${config.video.large.height}px`;
    card.style.width = `${config.video.large.width}px`;
    video.muted = false;
    if (cta) {
      videoWrapper.appendChild(cta);
    }
    if (bubble) {
      bubble.remove();
    }
  }

  videoWrapper.appendChild(video);
  videoWrapper.appendChild(progressBar);
  card.appendChild(videoWrapper);
  card.appendChild(button);
  if (bubble) {
    card.appendChild(bubble);
  }
  container.appendChild(card);
  video.play();
}
