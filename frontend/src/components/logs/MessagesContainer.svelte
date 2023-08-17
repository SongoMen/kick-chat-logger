<script lang="ts">
    import type { APIUserLog } from "$types/logs";
    import { padNumber } from "$lib/utils";
  
    export let label: string = "", username: string = "", channel: string = "";
    export let emotes: Record<string, string> = {};
    export let messages: APIUserLog[] = [];
    export let getMessages: (user: string, channel: string, period: string) => void;
    let hidden = messages.length === 0;

    const openCloseTab = async () => {
        if (!messages.length) {
            await getMessages(username, channel, label);
        }
        hidden = !hidden;
    };
  
    const formatMessageContent = (message: string) => {
      let finalMessage = '';
      for (const emote of message.split(' ')) {
        if (emotes[emote]) {
          finalMessage += `<img src="${emotes[emote]}" alt="${emote}" class="messages__emote">`;
        } else {
          finalMessage += emote;
        }
        finalMessage += ' ';
      }
      return finalMessage.trim();
    };
  
    const formatDate = (date: string) => {
      const d = new Date(date);
      return `${d.getFullYear()}-${padNumber(d.getMonth() + 1, 2)}-${padNumber(d.getDate(), 2)} ${padNumber(d.getHours(), 2)}:${padNumber(d.getMinutes(), 2)}:${padNumber(d.getSeconds(), 2)}`;
    };
  </script>
  
  <div class="messages" class:hidden>
    <h2 class="messages__label" on:click={openCloseTab}>{label}</h2>
    <ul class="messages__content">
      {#each messages as message}
        <li class="messages__text">
          <p>{formatDate(message.date)}</p>
          {#if (message.badges || []).length}
            <p>
              {#each message.badges as badge}
                <img src="/{badge}.svg" alt={badge} class="messages__badge">
              {/each}
            </p>
          {/if}
          <p>{username}:</p>
          <p>{@html formatMessageContent(message.message)}</p>
        </li>
      {/each}
    </ul>
  </div>
  
  <style lang="scss">
    .messages {
      display: flex;
      flex-direction: column;
      align-items: flex-start;
      width: 100%;
      border: 2px solid $bg-secondary;
      border-radius: 5px;
      padding: 15px;
      color: white;
      box-shadow: 0 0 10px rgba(0, 0, 0, 0.2);
      margin-bottom: 20px;
      overflow: auto;
      &.hidden {
        .messages {
            &__content {
                display: none;
            }
            &__label {
                border: 0;
                margin: 0;
            }
        }
      }

      &__content {
        margin: 0;
      }

      &__badge {
        margin-right: 5px;
        margin-top: 4px;
      }
  
      &__label {
        margin: 0;
        font-weight: bold;
        margin-bottom: 10px;
        width: 100%;
        cursor: pointer;
        border-bottom: 1px solid $grey;
      }
      :global(.messages__emote) {
        height: 20px;
        margin-left: 5px;
        margin-top: 5px;
      }
  
      ul {
        list-style: none;
        padding: 0;
  
        li {
          display: flex;
  
          p {
            margin: 0;
            margin-right: 5px;
            font-weight: 400;
            display: flex;
            align-items: flex-start;
  
            &:first-child {
              color: $grey-light;
              margin-right: 8px;
              width: 140px;
              min-width: 140px;
            }
  
            &:last-child {
              display: flex;
              flex-wrap: wrap;
            }
          }
        }
      }
    }
  </style>