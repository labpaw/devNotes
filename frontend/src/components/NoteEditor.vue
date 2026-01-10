<script setup>
const props = defineProps({
  modelValue: String,
  canDelete: Boolean,
  isOpen: Boolean,
});

const emit = defineEmits(['update:modelValue', 'delete', 'close']);

function onInput(event) {
  emit('update:modelValue', event.target.value);
}
</script>

<template>
  <Transition name="slide">
    <div v-if="isOpen" class="panel-overlay" @click.self="emit('close')">
      <div class="panel">
        <div class="panel-header">
          <button class="close-btn" @click="emit('close')">←</button>
          <button v-if="canDelete" class="delete-btn" @click="emit('delete')">Delete</button>
        </div>
        <textarea :value="modelValue" placeholder="Start typing your note..." @input="onInput" autofocus></textarea>
      </div>
    </div>
  </Transition>
</template>

<style scoped>
.panel-overlay {
  position: fixed;
  top: 0;
  left: 0;
  right: 0;
  bottom: 0;
  background: rgba(0, 0, 0, 0.3);
  z-index: 50;
}

.panel {
  position: fixed;
  left: 60px;
  top: 20px;
  bottom: 20px;
  width: 400px;
  background: #1e1e1e;
  border-radius: 12px;
  box-shadow: 0 20px 60px rgba(0, 0, 0, 0.5);
  display: flex;
  flex-direction: column;
  overflow: hidden;
}

.panel-header {
  display: flex;
  justify-content: space-between;
  align-items: center;
  padding: 12px 16px;
  background: #252525;
  border-bottom: 1px solid #333;
}

.close-btn {
  background: none;
  border: none;
  color: #888;
  font-size: 18px;
  cursor: pointer;
  padding: 4px 8px;
  border-radius: 4px;
}

.close-btn:hover {
  background: #333;
  color: #fff;
}

.delete-btn {
  padding: 6px 12px;
  background: transparent;
  color: #ef4444;
  border: 1px solid #ef4444;
  border-radius: 6px;
  cursor: pointer;
  font-size: 12px;
}

.delete-btn:hover {
  background: #ef4444;
  color: white;
}

textarea {
  flex: 1;
  padding: 20px;
  background: #1e1e1e;
  color: #e0e0e0;
  border: none;
  resize: none;
  font-size: 15px;
  line-height: 1.7;
  font-family: -apple-system, BlinkMacSystemFont, 'Segoe UI', sans-serif;
}

textarea:focus {
  outline: none;
}

textarea::placeholder {
  color: #555;
}

/* Slide transition */
.slide-enter-active,
.slide-leave-active {
  transition: all 0.25s ease;
}

.slide-enter-from,
.slide-leave-to {
  opacity: 0;
}

.slide-enter-from .panel,
.slide-leave-to .panel {
  transform: translateX(-20px);
  opacity: 0;
}
</style>
