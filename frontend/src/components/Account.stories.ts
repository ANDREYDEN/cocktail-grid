
import type { Meta, StoryObj } from '@storybook/vue3';

import ProfileImage from './ProfileImage.vue';
import { ref } from 'vue';
import { defaultAuthState, useAuth } from '../hooks/useAuth.mock';

const meta: Meta<typeof ProfileImage> = {
  component: ProfileImage,
  argTypes: {
  }
};

export default meta;

type Story = StoryObj<typeof ProfileImage>

export const Loading: Story = {
    async beforeEach() {
        useAuth.mockReturnValue({
            ...defaultAuthState,
            isLoading: ref(true),
            user: ref(undefined)
        })
    },
    render: (args) => ({
        components: { ProfileImage },
        setup() {
            return { args };
        },
        template: '<ProfileImage v-bind="args" />',
    }),
    args: {}
}

export const NoUser: Story = {
    async beforeEach() {
        useAuth.mockReturnValue({
            ...defaultAuthState,
            user: ref(undefined)
        })
    },
    render: (args) => ({
        components: { ProfileImage },
        setup() {
            return { args };
        },
        template: '<ProfileImage v-bind="args" />',
    }),
    args: {}
}

export const WithNoPicture: Story = {
    async beforeEach() {
        useAuth.mockReturnValue({
            ...defaultAuthState
        })
    },
    render: (args) => ({
        components: { ProfileImage },
        setup() {
            return { args };
        },
        template: '<ProfileImage v-bind="args" />',
    }),
    args: {}
}

export const WithBadPicture: Story = {
    async beforeEach() {
        useAuth.mockReturnValue({
            ...defaultAuthState,
            user: ref({
                picture: 'this is not a link'
            })
        })
    },
    render: (args) => ({
        components: { ProfileImage },
        setup() {
            return { args };
        },
        template: '<ProfileImage v-bind="args" />',
    }),
    args: {}
}

export const WithPicture: Story = {
    async beforeEach() {
        useAuth.mockReturnValue({
            ...defaultAuthState,
            user: ref({
                picture: 'https://picsum.photos/200'
            })
        })
    },
    render: (args) => ({
        components: { ProfileImage },
        setup() {
            return { args };
        },
        template: '<ProfileImage v-bind="args" />',
    }),
    args: {}
}