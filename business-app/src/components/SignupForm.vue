<template>
  <q-form class="text-center">
    <q-stepper class="signup" v-model="step" ref="stepper" animated>
      <q-step :name="1" title="Your location" icon="place" :done="step > 1">
        <div class="q-mb-md">{{ step }}/5</div>
        <h1 class="q-mb-lg">{{ $t('forms.signup.your_location_title') }}</h1>
        <q-select
          ref="countryField"
          v-model="formState.country"
          :options="getCountriesOptions()"
          class="q-mb-lg"
          :label="$t('forms.signup.country_field.label')"
          hide-bottom-space
          outlined
          emit-value
          map-options
          :loading="!getCountriesOptions()"
          :disable="!state.countries"
          :rules="[isValidCountry]"
          @update:model-value="formState.city = null"
        />
        <q-select
          v-if="isValidCountry() === true"
          v-model="formState.city"
          :options="getCitiesOptions()"
          class="q-mb-lg"
          :label="$t('forms.signup.city_field.label')"
          hide-bottom-space
          outlined
          emit-value
          map-options
          :loading="!getCitiesOptions()"
          :disable="!isValidCountry()"
          :rules="[isValidCity]"
        />
        <q-btn
          :disable="!isValidStep1()"
          class="full-width cta"
          icon-right="navigate_next"
          :label="$t('buttons.next')"
          hide-bottom-space
          @click="$refs.stepper.next()"
        />
      </q-step>

      <q-step
        :name="2"
        title="Your business details"
        icon="business"
        :done="step > 2"
      >
        <div class="q-mb-md">{{ step }}/5</div>
        <h1 class="q-mb-lg">{{ $t('forms.signup.your_business_details_title') }}</h1>
        <q-form>
          <q-select
            v-model="formState.businessType"
            :options="getBusinessTypesOptions()"
            class="q-mb-lg"
            :label="$t('forms.signup.business_type_field.label')"
            hide-bottom-space
            outlined
            emit-value
            map-options
            :disable="!getBusinessTypesOptions()"
            :loading="!state.businessTypes"
            :rules="[ isValidBusinessType ]"
          />
          <q-input
            v-model="formState.companyTitle"
            class="q-mb-lg"
            label="Company title"
            hide-bottom-space
            outlined
            required
            :rules="[
              (val) => val.length > 1 || 'Please provide at least 2 characters',
            ]"
          />
          <q-btn
            :disable="
              !formState.businessType || formState.companyTitle.length < 2
            "
            class="full-width cta"
            icon-right="navigate_next"
            :label="$t('buttons.next')"
            @click="$refs.stepper.next()"
          />
        </q-form>
      </q-step>

      <q-step
        :name="3"
        title="Your personal details"
        icon="account_circle"
        :done="step > 3"
      >
        <div class="q-mb-md">{{ step }}/5</div>
        <h1 class="q-mb-lg">Your personal details</h1>
        <q-form>
          <q-input
            v-model="formState.firstName"
            class="q-mb-lg"
            label="First name"
            hide-bottom-space
            outlined
            required
          />
          <q-input
            v-model="formState.lastName"
            class="q-mb-lg"
            label="Last name"
            hide-bottom-space
            outlined
          />
          <q-btn
            :disable="!formState.firstName || !formState.lastName"
            class="full-width cta"
            icon-right="navigate_next"
            label="Next"
            @click="$refs.stepper.next()"
          />
        </q-form>
      </q-step>

      <q-step
        :name="4"
        title="Your account details"
        icon="account_circle"
        :done="step > 4"
      >
        <div class="q-mb-md">{{ step }}/5</div>
        <h1 class="q-mb-lg">Your account details</h1>
        <q-form>
          <Vue3QTelInput
            ref="phoneField"
            v-model="formState.phone"
            class="q-mb-lg"
            label="Phone number"
            mask="###-###-###"
            unmasked-value
            hide-bottom-space
            outlined
            :rules="[isValidPhone]"
          />
          <q-input
            v-model="formState.email"
            class="q-mb-lg"
            label="Email"
            hide-bottom-space
            outlined
            :rules="[isValidEmail]"
          />
          <q-card class="important q-my-lg" flat>
            <q-card-section>
              <p>
                Phone number will be used to log in to your account. Please,
                make sure it belongs to you.
              </p>
              <p>
                We will send important service and account updates over email.
              </p>
            </q-card-section>
          </q-card>
          <q-btn
            :disable="!isValidStep4()"
            class="full-width cta"
            icon-right="navigate_next"
            label="Next"
            @click="$refs.stepper.next()"
          />
        </q-form>
      </q-step>

      <q-step :name="5" title="Choose subscription" icon="sell">
        <div class="q-mb-md">{{ step }}/5</div>
        <h1 class="q-mb-lg">Your subscription</h1>
        <p>
          <q-icon name="check" color="green" class="q-mr-sm" />Free for up
          to {{ freeSeatsLimit }} seats (employees)
        </p>
        <p>
          <q-icon name="check" color="green" class="q-mr-sm" />First
          {{ freeMonths }} months free for all
        </p>
        <p>
          <q-icon name="check" color="green" class="q-mr-sm" />${{
            seatPrice
          }}/month/seat for {{ freeSeatsLimit + 1 }} and more seats
          (employees)
        </p>
        <p class="q-mb-lg">
          <q-icon name="check" color="green" class="q-mr-sm" />You can add
          more seats later
        </p>

        <q-input
          type="number"
          v-model.number="formState.numberOfSeats"
          class="q-my-lg"
          label="Number of Seats"
          hide-bottom-space
          outlined
          :rules="[isValidNumberOfSeats]"
        />
        <q-card
          class="important q-mb-lg"
          flat
          v-if="isValidNumberOfSeats() === true"
        >
          <q-card-section v-if="formState.numberOfSeats <= freeSeatsLimit">
            Your subscription is <b>free</b>.
          </q-card-section>
          <q-card-section v-if="formState.numberOfSeats > freeSeatsLimit">
            You will pay
            <b>${{ formState.numberOfSeats * seatPrice }}/month</b
            >
            starting from
            {{
              getDateInThreeMonths().toLocaleDateString('en-US', {
                weekday: 'long',
                year: 'numeric',
                month: 'short',
                day: 'numeric',
              })
            }}
          </q-card-section>
        </q-card>
        <q-btn
          :disable="!isValidStep5()"
          class="full-width cta"
          label="Complete"
          @click="$refs.stepper.next()"
        />
      </q-step>
    </q-stepper>
    <div>
      <q-btn
        v-if="step > 1"
        class="link text-white q-mt-lg"
        @click="$refs.stepper.previous()"
        :label="$t('forms.signup.back_to_previous_step')"
        icon="navigate_before"
      />
    </div>
  </q-form>
</template>

<script lang="ts">
import { reactive, ref } from 'vue';
import { useI18n } from 'vue-i18n';
import { api } from 'boot/axios'
import Vue3QTelInput from 'vue3-q-tel-input'

export default {
  name: 'SignupForm',
  setup() {
    type SignupFormState = {
      country: string;
      city: string;
      businessType: string;
      companyTitle: string;
      firstName: string;
      lastName: string;
      phone: string;
      email: string;
      numberOfSeats: number;
    };

    const { t } = useI18n({ useScope: 'global' });

    const formState = reactive(<SignupFormState>(<unknown>{
      country: 'ge',
      city: 'ge.tbilisi',
      businessType: 'barber',
      companyTitle: 'Cool Barbers',
      firstName: 'Stanislav',
      lastName: 'Ivanov',
      phone: '571136842',
      email: 'resha.ru@gmail.com',
      numberOfSeats: null,
    }));
    const state = reactive({
      countries: [],
      businessTypes: [],
    });

    api.get('/countries')
      .then((response) => {
        state.countries = response.data;
      });

    api.get('/business-types')
      .then((response) => {
        state.businessTypes = response.data;
      });

    const businessTypes = reactive(['Barber', 'Dentist', 'Restaurant / Cafe']);

    const step = ref(1);

    const seatPrice = ref(5);
    const freeSeatsLimit = ref(2);
    const freeMonths = ref(3);

    return {
      step,
      formState,
      state,
      businessTypes,
      seatPrice,
      freeSeatsLimit,
      freeMonths,
      t,
      getCountriesOptions() {
        const options = [];
        for (const country of state.countries) {
          options.push({
            label: t('countries.' + country.name),
            value: country.name
          });
        }
        return options;
      },
      getCitiesOptions() {
        if (!state.countries || !formState.country) {
          return [];
        }
        const options = [];
        for (const country of state.countries) {
          if (country.name == formState.country) {
            for (const city of country.cities) {
              options.push({
                label: t('cities.' + city.name),
                value: city.name
              });
            }
          }
        }
        return options;
      },
      getBusinessTypesOptions() {
        const options = [];
        for (const businessType of state.businessTypes) {
          options.push({
            label: t('business_types.' + businessType.id),
            value: businessType.id
          });
        }
        return options;
      },
      getDateInThreeMonths() {
        let date = new Date();
        date.setMonth(date.getMonth() + freeMonths.value);
        return date;
      },
      isValidCountry() {
        return (
          formState.country != null || t('forms.signup.country_field.error')
        );
      },
      isValidCity() {
        return formState.city != null || t('forms.signup.city_field.error');
      },
      isValidBusinessType() {
        return formState.businessType != null || t('forms.signup.business_type_field.error');
      },
      isValidEmail() {
        const emailPattern =
          /^(?=[a-zA-Z0-9@._%+-]{6,254}$)[a-zA-Z0-9._%+-]{1,64}@(?:[a-zA-Z0-9-]{1,63}\.){1,8}[a-zA-Z]{2,63}$/;
        return emailPattern.test(formState.email) || 'Invalid email';
      },
      isValidPhone() {
        const phonePattern = /^[0-9]{9,10}$/;
        return phonePattern.test(formState.phone) || 'Invalid phone number';
      },
      isValidNumberOfSeats() {
        return (
          (formState.numberOfSeats > 0 && formState.numberOfSeats < 1000) ||
          'Please choose from 1 to 999 seats'
        );
      },
      isValidStep1() {
        return this.isValidCountry() === true && this.isValidCity() === true;
      },
      isValidStep4() {
        return this.isValidPhone() === true && this.isValidEmail() === true;
      },
      isValidStep5() {
        return this.isValidNumberOfSeats() === true;
      },
    };
  },
};
</script>

<style scoped></style>
