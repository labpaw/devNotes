<script setup>
import { ref, onMounted } from 'vue';
import { GetNotes, SaveNote, DeleteNote } from '../wailsjs/go/main/App';
import { WindowSetPosition, WindowSetSize, ScreenGetAll } from '../wailsjs/runtime/runtime';

const notes = ref([]);
const selectedId = ref(null);
const editContent = ref('');
const isExpanded = ref(false);

const colors = ['#3b82f6', '#10b981', '#f59e0b', '#ef4444', '#8b5cf6', '#ec4899'];

const COLLAPSED_WIDTH = 44;
const EXPANDED_WIDTH = 400;
const TAB_HEIGHT = 40;
const TAB_GAP = 4;
const PADDING = 16;

function calculateHeight() {
  // +1 for the add button
  const numTabs = notes.value.length + 1;
  return PADDING + numTabs * TAB_HEIGHT + (numTabs - 1) * TAB_GAP;
}

onMounted(async () => {
  await loadNotes();
  const screenWidth = window.screen.width;
  const height = calculateHeight();
  await WindowSetPosition(screenWidth - COLLAPSED_WIDTH, 100);
  await WindowSetSize(COLLAPSED_WIDTH, height);
});

async function loadNotes(skipResize = false) {
  notes.value = await GetNotes();
  // Only resize when collapsed, not during expansion
  if (!skipResize && !isExpanded.value) {
    const height = calculateHeight();
    const screenWidth = window.screen.width;
    await WindowSetSize(COLLAPSED_WIDTH, height);
    await WindowSetPosition(screenWidth - COLLAPSED_WIDTH, 100);
  }
}

function getColor(index) {
  return colors[index % colors.length];
}

function getPreview(text) {
  const firstLine = text.split('\n')[0] || 'New';
  return firstLine.length > 8 ? firstLine.slice(0, 8) : firstLine;
}

async function selectNote(note) {
  if (selectedId.value === note.id && isExpanded.value) {
    await collapse();
  } else {
    selectedId.value = note.id;
    editContent.value = note.content;
    await expand();
  }
}

async function expand() {
  isExpanded.value = true;
  const screenWidth = window.screen.width;
  // Instantly set size and position - no animation
  await WindowSetSize(EXPANDED_WIDTH, 400);
  await WindowSetPosition(screenWidth - EXPANDED_WIDTH, 100);
}

async function collapse() {
  isExpanded.value = false;
  selectedId.value = null;
  const screenWidth = window.screen.width;
  const height = calculateHeight();
  await WindowSetPosition(screenWidth - COLLAPSED_WIDTH, 100);
  await WindowSetSize(COLLAPSED_WIDTH, height);
}

async function saveNote() {
  if (!editContent.value.trim()) return;
  await SaveNote(selectedId.value || '', editContent.value);
  // Skip resize when saving while expanded - we're already at correct size
  await loadNotes(true);
}

async function createNew() {
  const saved = await SaveNote('', 'New note');
  // Skip resize - we're about to expand anyway
  await loadNotes(true);
  selectedId.value = saved.id;
  editContent.value = saved.content;
  await expand();
}

async function deleteNote() {
  if (!selectedId.value) return;
  await DeleteNote(selectedId.value);
  selectedId.value = null;
  editContent.value = '';
  await loadNotes();
  await collapse();
}
</script>

<template>
  <div class="app">
    <!-- Note content panel - full window when expanded -->
    <div v-if="isExpanded && selectedId" class="content-panel">
      <div class="panel-header" style="-webkit-app-region: drag">
        <button class="delete-btn" @click="deleteNote" style="-webkit-app-region: no-drag">🗑</button>
        <button class="close-btn" @click="collapse" style="-webkit-app-region: no-drag">✕</button>
      </div>
      <textarea v-model="editContent" @input="saveNote" placeholder="Write your note..."></textarea>
    </div>

    <!-- Tabs on RIGHT edge - only show when collapsed -->
    <div v-if="!isExpanded" class="tabs-column">
      <div class="tab add-tab" @click="createNew" title="New Note">
        <span>+</span>
      </div>
      <div
        v-for="(note, index) in notes"
        :key="note.id"
        :class="['tab', { active: selectedId === note.id }]"
        :style="{ background: getColor(index) }"
        @click="selectNote(note)"
        :title="note.content.split('\n')[0]"
      >
        <span class="tab-label">{{ getPreview(note.content) }}</span>
      </div>
    </div>
  </div>
</template>

<style>
* {
  margin: 0;
  padding: 0;
  box-sizing: border-box;
}

html,
body {
  background: transparent !important;
  overflow: hidden;
}

body {
  font-family: -apple-system, BlinkMacSystemFont, 'Segoe UI', sans-serif;
  color: #e0e0e0;
}

.app {
  display: flex;
  height: 100vh;
  width: 100vw;
  background: transparent;
  position: relative;
}

/* Tabs column - ALWAYS on RIGHT edge, bulletproof positioning */
.tabs-column {
  position: absolute;
  top: 0;
  right: 0;
  width: 44px;
  min-width: 44px;
  max-width: 44px;
  display: flex;
  flex-direction: column;
  gap: 4px;
  padding: 8px 0;
  background: transparent;
  align-items: flex-end;
  flex-shrink: 0;
}

.tab {
  width: 44px;
  height: 40px;
  border-radius: 6px 0 0 6px;
  cursor: pointer;
  display: flex;
  align-items: center;
  justify-content: center;
  transition: width 0.15s ease;
  box-shadow: -2px 2px 6px rgba(0, 0, 0, 0.25);
  margin-right: 0 !important; /* Force flush with right edge */
  position: relative;
  right: 0;
}

.tab:hover {
  width: 60px;
}

.tab.active {
  width: 52px;
}

.tab-label {
  font-size: 9px;
  color: white;
  font-weight: 600;
  writing-mode: vertical-rl;
  text-orientation: mixed;
  text-shadow: 0 1px 2px rgba(0, 0, 0, 0.4);
}

.add-tab {
  background: #444;
}

.add-tab:hover {
  background: #3b82f6;
}

.add-tab span {
  font-size: 24px;
  color: white;
  font-weight: 300;
}

/* Content panel - full window when note is open */
.content-panel {
  flex: 1;
  margin: 10px;
  background: rgba(26, 26, 26, 0.98);
  border-radius: 12px;
  display: flex;
  flex-direction: column;
  overflow: hidden;
  box-shadow: 0 8px 32px rgba(0, 0, 0, 0.4);
}

.panel-header {
  display: flex;
  justify-content: space-between;
  align-items: center;
  padding: 12px 14px;
  background: #222;
  border-bottom: 1px solid #333;
}

.close-btn,
.delete-btn {
  background: rgba(255, 255, 255, 0.1);
  border: none;
  color: #888;
  font-size: 14px;
  cursor: pointer;
  width: 32px;
  height: 32px;
  border-radius: 6px;
  display: flex;
  align-items: center;
  justify-content: center;
}

.close-btn:hover {
  background: rgba(255, 255, 255, 0.15);
  color: #fff;
}

.delete-btn:hover {
  background: rgba(239, 68, 68, 0.2);
  color: #ef4444;
}

.content-panel textarea {
  flex: 1;
  padding: 16px;
  background: transparent;
  border: none;
  color: #e0e0e0;
  font-size: 14px;
  line-height: 1.6;
  resize: none;
  font-family: 'SF Mono', Monaco, 'Courier New', monospace;
}

.content-panel textarea:focus {
  outline: none;
}

.content-panel textarea::placeholder {
  color: #555;
}
</style>
