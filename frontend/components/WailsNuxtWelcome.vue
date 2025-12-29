<script setup lang="ts">
import nuxtLogo from "~/assets/images/nuxt.png";
import wailsLogo from "~/assets/images/logo-universal.png";
import { Greet } from "~~/bindings/changeme/internal/services/greetservice";
import { Events } from "@wailsio/runtime";

const name = ref<string>("");
const greeting = ref<string>("\xa0");

const greet = async () => {
  greeting.value = await Greet(name.value);
}

const quit = () => {
  Events.Emit("application-exit");
}

</script>
<template>
  <UContainer class="h-screen flex items-center justify-center">
    <UCard class="w-full max-w-sm">
      <template #header>
        <div class="flex items-center justify-around h-24 py-4">
          <img :src="wailsLogo" class="h-full object-contain" alt="Wails Logo" />
          <img :src="nuxtLogo" class="h-1/2 object-contain" alt="Nuxt Logo" />
        </div>
      </template>

      <div class="space-y-6">
        <div class="text-center">
          <h1 class="text-2xl font-bold mb-2">Welcome to Wails + Nuxt</h1>
          <p class="text-sm text-gray-500 dark:text-gray-400 min-h-[1.25rem]">
            {{ greeting }}
          </p>
        </div>

        <div class="flex gap-2">
          <UInput
            v-model="name"
            placeholder="Your name..."
            class="flex-1"
            @keyup.enter="greet"
          />
          <UButton @click="greet">
            Greet
          </UButton>
        </div>

        <UButton
          color="neutral"
          variant="subtle"
          block
          @click="quit"
        >
          Close App
        </UButton>
      </div>
    </UCard>
  </UContainer>
</template>

<style scoped>
/* No additional styles needed as we use Nuxt UI and Tailwind */
</style>

