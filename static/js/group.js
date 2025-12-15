

/**
 * @fileoverview This script handles the client-side logic for a group initiative tracker in a tabletop RPG application.
 * It manages WebSocket communication for real-time updates, renders the initiative list, handles user interactions
 * for adding players/monsters, managing HP, conditions, and turn progression. It supports DM (Dungeon Master) controls
 * like editing entities, reordering via drag-and-drop, and resetting the initiative.
 */

/**
 * Immediately Invoked Function Expression (IIFE) to encapsulate the initiative tracker logic.
 * Sets up WebSocket connection, event listeners, and UI rendering for the group initiative system.
 */
(() => {
  // ... (code omitted)

  /**
   * Handles incoming WebSocket messages, specifically updating the UI state based on 'state' messages.
   * @param {MessageEvent} ev - The WebSocket message event containing JSON data.
   */
  ws.addEventListener('message', (ev) => {
    // ... (code omitted)
  });

  /**
   * Renders the list of entities (players and monsters) in the initiative order.
   * Creates DOM elements for each entity, displaying name, initiative, conditions, HP (for monsters if DM),
   * and controls (if DM). Highlights the current turn.
   * @param {Array} entries - Array of entity objects with properties like id, name, initiative, bonus, tags, type, hp, maxHp.
   * @param {number} turn - Index of the current turn in the entries array.
   */
  function render(entries, turn) {
    // ... (code omitted)
  }

  /**
   * Sends a message to the WebSocket server if the connection is open.
   * @param {string} type - The message type (e.g., 'addPlayer', 'damage').
   * @param {Object} data - The data payload for the message.
   */
  function wsSend(type, data) {
    // ... (code omitted)
  }

  /**
   * Escapes HTML characters in a string to prevent XSS.
   * @param {string} s - The string to escape.
   * @returns {string} The escaped string.
   */
  function escapeHtml(s){return s.replace(/[&<>"]+/g, c=>({'&':'&amp;','<':'&lt;','>':'&gt;','"':'&quot;'}[c]));}

  /**
   * Calculates the percentage of a value relative to a maximum, clamped between 0 and 100.
   * @param {number} a - The current value.
   * @param {number} b - The maximum value.
   * @returns {number} The percentage as an integer.
   */
  function pct(a,b){ if(!b) return 0; return Math.max(0,Math.min(100, Math.round(a*100/b))); }

  // Forms
  // ... (event listeners for forms)

  /**
   * Event listener for the player form submission. Adds a new player via WebSocket.
   */
  pf?.addEventListener('submit', (e) => {
    // ... (code omitted)
  });

  /**
   * Event listener for the roll button. Adds a player with rolled initiative via WebSocket.
   */
  document.getElementById('rollBtn')?.addEventListener('click', ()=>{
    // ... (code omitted)
  });

  /**
   * Event listener for the monster form submission. Adds a new monster via WebSocket.
   */
  mf?.addEventListener('submit', (e) => {
    // ... (code omitted)
  });

  /**
   * Event listener for the next button. Advances to the next turn via WebSocket.
   */
  nextBtn?.addEventListener('click', ()=> wsSend('next', {}));

  /**
   * Event listener for the reset button (DM only). Resets the initiative, removing all entities.
   */
  resetBtn?.addEventListener('click', () => {
    // ... (code omitted)
  });

  // Drag and drop for DM
  // ... (Sortable setup)
})();