<script setup>
defineProps({
  notes: Array,
  selectedId: String,
});

const emit = defineEmits(['select', 'new']);

function getPreview(text) {
  const firstLine = text.split('\n')[0] || 'New';
  return firstLine.length > 15 ? firstLine.slice(0, 15) : firstLine;
}

function getColor(index) {
  const colors = ['#3b82f6', '#10b981', '#f59e0b', '#ef4444', '#8b5cf6', '#ec4899'];
  return colors[index % colors.length];
}
</script>

<template>
  <div class="tabs-container">
    <div class="tab new-tab" @click="emit('new')" title="New Note">
      <span class="tab-icon">+</span>
    </div>
    <div
      v-for="(note, index) in notes"
      :key="note.id"
      :class="['tab', { active: selectedId === note.id }]"
      :style="{ borderLeftColor: getColor(index) }"
      @click="emit('select', note)"
      :title="note.content.split('\n')[0]"
    >
      <span class="tab-text">{{ getPreview(note.content) }}</span>
    </div>
  </div>
</template>

<style scoped>
.tabs-container {
  position: fixed;
  left: 0;
  top: 50%;
  transform: translateY(-50%);
  display: flex;
  flex-direction: column;
  gap: 4px;
  z-index: 100;
}

.tab {
  width: 40px;
  height: 48px;
  background: #2a2a2a;
  border-left: 4px solid #555;
  border-radius: 0 8px 8px 0;
  cursor: pointer;
  display: flex;
  align-items: center;
  justify-content: center;
  transition: all 0.2s ease;
  overflow: hidden;
  writing-mode: vertical-rl;
  text-orientation: mixed;
}

.tab:hover {
  width: 120px;
  background: #333;
}

.tab.active {
  width: 50px;
  background: #3a3a3a;
}

.tab.active:hover {
  width: 120px;
}

.tab-text {
  font-size: 11px;
  color: #ccc;
  white-space: nowrap;
  overflow: hidden;
  text-overflow: ellipsis;
  transform: rotate(180deg);
}

.tab:hover .tab-text {
  writing-mode: horizontal-tb;
  transform: none;
  padding: 0 8px;
}

.new-tab {
  background: #3b82f6;
  border-left-color: #3b82f6;
  width: 40px;
}

.new-tab:hover {
  width: 40px;
  background: #2563eb;
}

.tab-icon {
  font-size: 20px;
  font-weight: bold;
  color: white;
  writing-mode: horizontal-tb;
}
</style>
