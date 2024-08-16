import type { Meta, StoryObj } from '@storybook/vue3';

import CustomButton from './CustomButton.vue';

const meta: Meta<typeof CustomButton> = {
  component: CustomButton,
  argTypes: {
    iconPosition: { control: 'select', options: ['left', 'right']}
  }
};

export default meta;

type Story = StoryObj<typeof CustomButton>

export const Outlined: Story = {
    render: (args) => ({
        components: { CustomButton },
        setup() {
            return { args };
        },
        template: '<CustomButton v-bind="args">My Text</CustomButton>',
    }),
    args: {
        outlined: true,
    }
}

export const Text: Story = {
    args: {
        outlined: false
    },
    render: args => ({
        components: {
            CustomButton
        },

        setup() {
            return { args };
        },

        template: "<CustomButton v-bind=\"args\">My Text</CustomButton>"
    })
};

export const Loading: Story = {
    args: {
        outlined: true,
        loading: true
    },
    render: args => ({
        components: {
            CustomButton
        },

        setup() {
            return { args };
        },

        template: "<CustomButton v-bind=\"args\">My Text</CustomButton>"
    })
};